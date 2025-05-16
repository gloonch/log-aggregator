package service

import (
	"math/rand"
	"time"

	"github.com/gloonch/log-aggregator/producer-market/internal/domain"
)

func GenerateDailyCandles(symbol string, startDate, endDate time.Time) []domain.Candle {
	var candles []domain.Candle
	rand.Seed(time.Now().UnixNano())

	price := 1900.0 + rand.Float64()*100 // base price around gold range

	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		open := price
		high := open + rand.Float64()*20
		low := open - rand.Float64()*20
		if low < 0 {
			low = 0.1 // so we don't have negative prices
		}
		close := low + rand.Float64()*(high-low)

		candles = append(candles, domain.Candle{
			Symbol:    symbol,
			Open:      round(open),
			High:      round(high),
			Low:       round(low),
			Close:     round(close),
			StartTime: d.Format(time.RFC3339),
			EndTime:   d.AddDate(0, 0, 1).Format(time.RFC3339),
		})

		price = close // next day's open price
	}

	return candles
}

func round(val float64) float64 {
	return float64(int(val*100)) / 100
}
