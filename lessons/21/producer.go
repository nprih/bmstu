package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

const (
	kafkaAddress = "localhost:9092"
	topic        = "my-topic"
)

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{kafkaAddress},
		Topic:   topic,
	})
	defer w.Close()
	messages := []string{
		"Hello",
		"my friends",
		"I am",
		"Dmitry",
	}
	sendMessagesToKafka(w, messages)
	log.Println("All messages were sent")
}

func sendMessagesToKafka(w *kafka.Writer, messages []string) {
	for i, msg := range messages {
		newMessage := kafka.Message{
			Key:   []byte(fmt.Sprintf("Message #%d", i+1)),
			Value: []byte(msg),
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		err := w.WriteMessages(ctx, newMessage)
		if err != nil {
			fmt.Println("error with kafka messages:", err)
		} else {
			fmt.Printf("Message #%d was sent\n", i)
		}
		time.Sleep(4 * time.Second)
	}
}
