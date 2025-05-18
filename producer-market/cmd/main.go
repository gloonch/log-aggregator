package main

import (
	"github.com/gloonch/log-aggregator/producer-market/internal/domain"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gloonch/log-aggregator/producer-market/internal/kafka"
	"github.com/gloonch/log-aggregator/producer-market/internal/service"
)

func main() {
	broker := os.Getenv("BROKER_ADDR")
	if broker == "" {
		broker = "kafka:9092"
	}
	symbol := os.Getenv("SYMBOL")
	if symbol == "" {
		symbol = "XAUUSD"
	}

	timeframe := os.Getenv("TIMEFRAME")
	if timeframe == "" {
		timeframe = "1h"
	}

	topic := strings.ToLower("market.candles")

	// TODO: this has to be taken from the API
	start := time.Now().UTC().AddDate(0, 0, -5) // last 5 days
	end := time.Now().UTC()

	conn := kafka.NewConnection("tcp", broker, topic)
	defer conn.Close()

	writer := kafka.NewCandleWriter(broker, topic)
	defer writer.Close()

	var candles []domain.Candle

	switch timeframe {
	case "daily":
		candles = service.GenerateDailyCandles(symbol, start, end)
	case "1h":
		candles = service.GenerateHourlyCandles(symbol, start, end)
	case "4h":
		candles = service.GenerateFourHoursCandles(symbol, start, end)
	default:
		log.Fatalf("❌ unsupported timeframe: %s", timeframe)
	}

	for _, c := range candles {
		err := writer.Write(c)
		if err != nil {
			log.Printf("❌ failed to write candle: %v", err)
		} else {
			log.Printf("✅ candle sent: %+v", c)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
