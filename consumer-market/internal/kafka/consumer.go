package kafka

import (
	"context"
	"encoding/json"
	domain "github.com/gloonch/log-aggregator/consumer-market/internal/model"
	"github.com/gloonch/log-aggregator/consumer-market/internal/store"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type CandleReader struct {
	reader *kafka.Reader
}

func StartConsumer(ctx context.Context, broker, topic, groupID, matchKey string, store *store.RedisStore) {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 1e3,
		MaxBytes: 10e6,
	})

	log.Printf("👂 [%s] waiting for messages with key = %s", groupID, matchKey)

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("❌ [%s] read error: %v", groupID, err)

			continue
		}

		//log.Printf("Raw Key: %s", string(msg.Key))

		if string(msg.Key) != matchKey {
			continue // skip other timeframe messages
		}

		var candle domain.Candle
		if err := json.Unmarshal(msg.Value, &candle); err != nil {
			log.Printf("❌ [%s] failed to unmarshal: %v", groupID, err)

			continue
		}

		log.Printf("✅ [%s] received: %+v", groupID, candle)

		candleJSON, _ := json.Marshal(candle) // could also use `msg.Value` but this is more accurate
		ts, _ := time.Parse(time.RFC3339, candle.StartTime)

		err = store.SaveCandle(candle.Symbol, matchKey, candleJSON, ts.Unix())
		if err != nil {
			log.Printf("❌ [%s] failed to save to Redis: %v", groupID, err)
		} else {
			log.Printf("💾 [%s] saved to Redis: %s [%s]", groupID, candle.Symbol, matchKey)
		}
	}
}
