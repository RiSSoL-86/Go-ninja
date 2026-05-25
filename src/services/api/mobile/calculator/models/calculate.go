package models

import (
	"fmt"

	"github.com/danielgtaylor/huma/v2"
)

type CalculateRequest struct {
	Args     []float64 `json:"args" minItems:"1" doc:"Numbers to calculate"`
	Operator string    `json:"operator" enum:"+,-,*,/" doc:"Arithmetic operation"`
}

func (r *CalculateRequest) Resolve(_ huma.Context) []error {
	if r.Operator != "/" {
		return nil
	}

	for i := 1; i < len(r.Args); i++ {
		if r.Args[i] == 0 {
			return []error{&huma.ErrorDetail{
				Location: fmt.Sprintf("body.args[%d]", i),
				Message:  "must not be zero when operator is /",
				Value:    r.Args[i],
			}}
		}
	}

	return nil
}

type CalculateResponse struct {
	ID     int64   `json:"id,omitempty"`
	Result float64 `json:"result"`
}

type CalculateInput struct {
	Body CalculateRequest
}

type CalculateOutput struct {
	Body CalculateResponse
}
