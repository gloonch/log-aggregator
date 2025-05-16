package domain

type Candle struct {
	Symbol    string  `json:"symbol"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	StartTime string  `json:"startTime"`
	EndTime   string  `json:"endTime"`
}
