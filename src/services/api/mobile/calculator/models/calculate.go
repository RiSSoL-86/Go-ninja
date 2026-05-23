package models

type CalculateRequest struct {
	Args     []float64 `json:"args"`
	Operator string    `json:"operator"`
}

type CalculateResponse struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}
