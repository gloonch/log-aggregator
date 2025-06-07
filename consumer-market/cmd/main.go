package main

import (
	"context"
	"github.com/gloonch/log-aggregator/consumer-market/internal/kafka"
	"github.com/gloonch/log-aggregator/consumer-market/internal/store"
	"os"
)

func main() {

	broker := os.Getenv("BROKER_ADDR")
	if broker == "" {
		broker = "kafka:9092"
	}

	topic := "market.candles"

	ctx := context.Background()

	redisStore := store.NewRedisStore("redis:6379")

	go kafka.StartConsumer(ctx, broker, topic, "consumer-gold", "GOLD", redisStore)
	go kafka.StartConsumer(ctx, broker, topic, "consumer-silver", "SILVER", redisStore)
	go kafka.StartConsumer(ctx, broker, topic, "consumer-iron", "IRON", redisStore)

	select {}
}
