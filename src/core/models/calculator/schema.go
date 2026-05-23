package calculator

import "time"

type CalculationRecord struct {
	ID        int64     `json:"id"`
	Args      []float64 `json:"args"`
	Operator  string    `json:"operator"`
	Result    float64   `json:"result"`
	CreatedAt time.Time `json:"created_at"`
}
