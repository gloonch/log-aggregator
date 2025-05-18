package service

import (
	"math/rand"
	"time"

	"github.com/gloonch/log-aggregator/producer-market/internal/domain"
)

// GenerateCandles returns OHLC candles with arbitrary duration (in minutes)
func GenerateCandles(symbol string, startTime, endTime time.Time, intervalMinutes int) []domain.Candle {
	var candles []domain.Candle
	rand.Seed(time.Now().UnixNano())

	price := 3000.0 + rand.Float64()*100 // base price for gold

	interval := time.Duration(intervalMinutes) * time.Minute

	for t := startTime; !t.After(endTime); t = t.Add(interval) {
		open := price
		high := open + rand.Float64()*10
		low := open - rand.Float64()*10
		if low < 0 {
			low = 0.1
		}
		close := low + rand.Float64()*(high-low)

		candles = append(candles, domain.Candle{
			Symbol:    symbol,
			Open:      round(open),
			High:      round(high),
			Low:       round(low),
			Close:     round(close),
			StartTime: t.Format(time.RFC3339),
			EndTime:   t.Add(interval).Format(time.RFC3339),
		})

		price = close
	}

	return candles
}

// Convenience wrappers for common timeframes
func GenerateDailyCandles(symbol string, start, end time.Time) []domain.Candle {
	return GenerateCandles(symbol, start, end, 1440)
}

func GenerateHourlyCandles(symbol string, start, end time.Time) []domain.Candle {
	return GenerateCandles(symbol, start, end, 60)
}

func GenerateFourHoursCandles(symbol string, start, end time.Time) []domain.Candle {
	return GenerateCandles(symbol, start, end, 240)
}

func round(val float64) float64 {
	return float64(int(val*100)) / 100
}
