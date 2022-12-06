1. Create a topic:

```
kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 3 --partitions 6 --topic weapons
```


2. Publish some test data to the topic:
```
kafka-console-producer --broker-list localhost:9092 --topic weapons

hero: "Loken"", weapon: "Boltgun"
hero: "Qruze"", weapon: "Las Rifle"
hero: "Garro"", weapon: "Krak Grenade"
```

3. Consume data from the topic:
```
kafka-console-consumer --bootstrap-server localhost:9092 --topic weapons --from-beginning

hero: "Loken"", weapon: "Boltgun"
hero: "Qruze"", weapon: "Las Rifle"
hero: "Garro"", weapon: "Krak Grenade"
```