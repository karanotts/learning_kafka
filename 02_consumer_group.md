Group 1, Consumer 1

```
kafka-console-consumer --bootstrap-server localhost:9092 --topic weapons --group 1 > group1_consumer1.txt
```


Group 2, Consumer 1
```
kafka-console-consumer --bootstrap-server localhost:9092 --topic weapons --group 2 > group2_consumer1.txt
```


Group 2, Consumer 2
```
kafka-console-consumer --bootstrap-server localhost:9092 --topic weapons --group 2 > group2_consumer2.txt
```