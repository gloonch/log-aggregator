package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "market.ticks",
		GroupID: "log-aggregator-group",
	})
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}
		log.Printf("received:, %s!", string(msg.Value))
	}
}
