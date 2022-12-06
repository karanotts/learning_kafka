Create a topic:
```
kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 3 --partitions 6 --topic weapons
```
Setup basic config `kafka-go.properties`
```
bootstrap.servers=localhost:9092
acks=all
```
Initialise Go module and download Confluent Kafka GO
```
go mod init kafka-go
go get github.com/confluentinc/confluent-kafka-go/kafka
```
Compile the producer
```
go build -o out/producer util.go producer.go
```
Compile the consumer
```
go build -o out/consumer util.go consumer.go
```
Produce events
```
./out/producer go-kafka.properties
```
Consume events
```
./out/consumer go-kafka.properties
```
Rejoice seeing that Garro got a chainsword!
```
./out/consumer kafka-go.properties 

Consumed event from topic weapons: key = Voyen      value = Stub Rifle
Consumed event from topic weapons: key = Voyen      value = Krak Grenade
Consumed event from topic weapons: key = Tarvitz    value = Boltgun
Consumed event from topic weapons: key = Tarvitz    value = Lasgun
Consumed event from topic weapons: key = Qruze      value = Lasgun
Consumed event from topic weapons: key = Garro      value = Chainsword
Consumed event from topic weapons: key = Qruze      value = Krak Grenade
Consumed event from topic weapons: key = Voyen      value = Chainsword
Consumed event from topic weapons: key = Loken      value = Krak Grenade
Consumed event from topic weapons: key = Temeter    value = Lasgun
```