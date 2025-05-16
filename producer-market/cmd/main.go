package main

import (
	"log"
	"os"
	"time"

	"github.com/gloonch/log-aggregator/producer-market/internal/kafka"
	"github.com/gloonch/log-aggregator/producer-market/internal/service"
)

func main() {
	broker := os.Getenv("BROKER_ADDR")
	if broker == "" {
		broker = "localhost:9092"
	}

	topic := "gold.candles.daily"
	symbol := "GOLD"

	// تعیین بازه زمانی تولید داده
	start := time.Date(2025, 5, 10, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 5, 14, 0, 0, 0, 0, time.UTC)

	// ایجاد اتصال و ساخت topic
	conn := kafka.NewConnection("tcp", broker, topic)
	defer conn.Close()

	writer := kafka.NewCandleWriter(broker, topic)
	defer writer.Close()

	candles := service.GenerateDailyCandles(symbol, start, end)

	for _, c := range candles {
		err := writer.Write(c)
		if err != nil {
			log.Printf("❌ failed to write candle: %v", err)
		} else {
			log.Printf("✅ candle sent: %+v", c)
		}
		time.Sleep(500 * time.Millisecond) // برای نمایش بهتر لاگ‌ها
	}
}
