package main

import (
	"context"
	"github.com/gloonch/log-aggregator/consumer-market/internal/api"
	"github.com/gloonch/log-aggregator/consumer-market/internal/kafka"
	pb "github.com/gloonch/log-aggregator/consumer-market/internal/proto/github.com/gloonch/log-aggregator/api/candlepb"
	"github.com/gloonch/log-aggregator/consumer-market/internal/store"
	"google.golang.org/grpc"
	"log"
	"net"
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

	lis, _ := net.Listen("tcp", ":50051")
	grpcServer := grpc.NewServer()

	pb.RegisterCandleServiceServer(grpcServer, api.NewCandleServer(redisStore))

	log.Println("🚀 gRPC server started on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	go kafka.StartConsumer(ctx, broker, topic, "consumer-gold", "GOLD", redisStore)
	go kafka.StartConsumer(ctx, broker, topic, "consumer-silver", "SILVER", redisStore)
	go kafka.StartConsumer(ctx, broker, topic, "consumer-iron", "IRON", redisStore)

	select {}
}
