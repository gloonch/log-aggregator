package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	// STEP 1: Create Topic with AdminClient
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		log.Fatalf("failed to connect to Kafka: %v", err)
	}
	defer conn.Close()

	topic := "logs"
	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     3,
		ReplicationFactor: 1,
	})
	if err != nil {
		log.Printf("could not create topic (might already exist): %v", err)
	} else {
		log.Printf("topic %s created or already exists", topic)
	}

	// STEP 2: Set up Kafka Writer
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	// STEP 3: Send Sample Messages
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("log message %d", i)
		err := writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(fmt.Sprintf("key-%d", i)),
			Value: []byte(msg),
		})
		if err != nil {
			log.Printf("failed to write message: %v", err)
		} else {
			log.Printf("sent: %s", msg)
		}
		time.Sleep(1 * time.Second)
	}
}
