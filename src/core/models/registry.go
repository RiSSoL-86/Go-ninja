package models

import (
	"app/src/core/models/calculator"
)

func GetModels() []any {
	return []any{
		&calculator.CalculationRecord{},
		// Add new models here
	}
}
