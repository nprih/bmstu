package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	kafkaAddress = "localhost:9092"
	topic        = "my-topic"
	group        = "MyGroup"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaAddress},
		Topic:   topic,
		GroupID: group,
	})
	defer r.Close()
	log.Println("Consumer connected")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	readMessagesFromKafka(ctx, r)
}
func readMessagesFromKafka(ctx context.Context, r *kafka.Reader) {
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("I got message: topic: %s, key:%s, value:%s\n", msg.Topic, string(msg.Key), string(msg.Value))
	}
}
