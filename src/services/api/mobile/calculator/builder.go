package calculator

import (
	"app/src/services/api/mobile/calculator/services"
)

func BuildCalculator() *Handler {
	repository := NewMemoryRepository()

	calculateServices := services.NewCalculateService(repository)
	historyServices := services.NewHistoryService(repository)

	return NewHandler(calculateServices, historyServices)
}
