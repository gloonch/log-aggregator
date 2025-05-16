package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gloonch/log-aggregator/producer-market/internal/domain"
	"github.com/segmentio/kafka-go"
)

type CandleWriter struct {
	writer *kafka.Writer
}

// NewConnection creates a TCP connection and ensures topic exists
func NewConnection(network, address, topic string) *kafka.Conn {
	conn, err := kafka.Dial(network, address)
	if err != nil {
		log.Fatalf("failed to connect to kafka: %v", err)
	}

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})

	log.Println("Creating topic:", topic)

	if err != nil {
		log.Printf("Failed to create topic %s: %v", topic, err)
	} else {
		log.Printf("Topic %s created or already exists", topic)
	}

	return conn
}

// NewCandleWriter creates a Kafka writer for Candle messages
func NewCandleWriter(broker string, topic string) *CandleWriter {
	return &CandleWriter{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(broker),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

// Write publishes a Candle message to Kafka
func (cw *CandleWriter) Write(candle domain.Candle) error {
	data, err := json.Marshal(candle)
	if err != nil {
		return err
	}
	return cw.writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Key:   []byte(candle.Symbol),
			Value: data,
		})
}

// Close shuts down the Kafka writer
func (cw *CandleWriter) Close() {
	_ = cw.writer.Close()
}
