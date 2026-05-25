package calculator

import (
	"app/src/services/api/mobile/calculator/services"

	"gorm.io/gorm"
)

type Dependencies struct {
	handler *Handler
}

func NewDependencies(db *gorm.DB) *Dependencies {
	calculatorRepository := NewRepository(db)

	return &Dependencies{
		handler: NewHandler(
			services.NewCalculateService(calculatorRepository),
			services.NewHistoryService(calculatorRepository),
		),
	}
}
