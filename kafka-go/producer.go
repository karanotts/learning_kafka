package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <config-file-path\n", os.Args[0])
		os.Exit(1)
	}

	configFile := os.Args[1]
	conf := ReadConfig(configFile)

	topic := "weapons"
	p, err := kafka.NewProducer(&conf)

	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Goroutine to handle message delivery reports and possibly other event types (errs, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	heroes := [...]string{"Loken", "Tarvitz", "Garro", "Temeter", "Qruze", "Voyen"}
	weapons := [...]string{"Lasgun", "Chainsword", "Stub Rifle", "Boltgun", "Krak Grenade"}

	for n := 0; n < 10; n++ {
		key := heroes[rand.Intn(len(heroes))]
		data := weapons[rand.Intn(len(weapons))]

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(data),
		}, nil)
	}

	// Wait for all messages to be delivered
	p.Flush(15 * 1000)
	p.Close()
}
