package main

import (
	"context"
	"github.com/gloonch/log-aggregator/consumer-market/internal/kafka"
	"os"
)

func main() {

	broker := os.Getenv("BROKER_ADDR")
	if broker == "" {
		broker = "kafka:9092"
	}

	topic := "market.candles"

	ctx := context.Background()

	go kafka.StartConsumer(ctx, broker, topic, "consumer-gold", "GOLD")
	go kafka.StartConsumer(ctx, broker, topic, "consumer-silver", "SILVER")
	go kafka.StartConsumer(ctx, broker, topic, "consumer-iron", "IRON")

	select {}
}
