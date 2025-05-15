package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"math/rand"
	"time"
)

type Tick struct {
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	TimeStamp string  `json:"timestamp"`
}

func main() {

	rand.Seed(time.Now().UnixNano())

	conn, err := kafka.Dial("tcp", "kafka:9092")
	if err != nil {
		log.Fatalf("failed to connect to Kafka: %v", err)
	}
	defer conn.Close()

	topic := "market.ticks"
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

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	symbols := []string{"SOXL", "XRP", "AAPL"}

	for {
		tick := Tick{
			Symbol:    symbols[rand.Intn(len(symbols))],
			Price:     10 + rand.Float64()*100,
			TimeStamp: time.Now().UTC().Format(time.RFC3339),
		}

		data, _ := json.Marshal(tick)

		err := writer.WriteMessages(
			context.Background(),
			kafka.Message{
				Key:   []byte(tick.Symbol),
				Value: data,
			},
		)
		if err != nil {
			log.Printf("failed to send tick: %v", err)
		} else {
			log.Printf("sent tick: %s", data)
		}
		time.Sleep(1 * time.Second)
	}
}
