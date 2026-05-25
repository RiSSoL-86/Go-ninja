package models

import "time"

type HistoryItem struct {
	ID        int64     `json:"id"`
	Args      []float64 `json:"args"`
	Operator  string    `json:"operator"`
	Result    float64   `json:"result"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}

type HistoryOutput struct {
	Body []HistoryItem
}
