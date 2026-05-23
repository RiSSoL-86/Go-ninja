package calculator

import (
	"app/src/services/api/mobile/calculator/services"
)

func BuildCalculator() *Handler {
	repository := NewMemoryRepository()

	calculateService := services.NewCalculateService(repository)
	historyService := services.NewHistoryService(repository)

	return NewHandler(calculateService, historyService)
}
