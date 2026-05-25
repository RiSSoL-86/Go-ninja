package mobile

import (
	"app/src/services/api/mobile/calculator"

	"gorm.io/gorm"
)

type Dependencies struct {
	Calculator *calculator.Dependencies
}

func NewDependencies(db *gorm.DB) *Dependencies {
	return &Dependencies{
		Calculator: calculator.NewDependencies(db),
	}
}
