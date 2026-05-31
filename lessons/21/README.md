go get github.com/segmentio/kafka-go

docker exec -it kafka /opt/kafka/bin/kafka-topics.sh \
--create \
--topic my-topic \
--bootstrap-server localhost:9092 \
--partitions 3 \
--replication-factor 1