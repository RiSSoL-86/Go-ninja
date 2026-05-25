package calculator

import (
	calcModels "app/src/services/api/mobile/calculator/models"
	"context"
	"log"

	"github.com/danielgtaylor/huma/v2"
)

type ICalculateService interface {
	Execute(req calcModels.CalculateRequest) (calcModels.CalculateResponse, error)
}

type IHistoryService interface {
	GetHistory() ([]calcModels.HistoryItem, error)
}

type Handler struct {
	calculateService ICalculateService
	historyService   IHistoryService
}

func NewHandler(calculateService ICalculateService, historyService IHistoryService) *Handler {
	return &Handler{
		calculateService: calculateService,
		historyService:   historyService,
	}
}

func (h *Handler) Calculate(_ context.Context, input *calcModels.CalculateInput) (*calcModels.CalculateOutput, error) {
	res, err := h.calculateService.Execute(input.Body)
	if err != nil {
		log.Printf("Calculate error: %v", err)
		return nil, huma.Error500InternalServerError("failed to calculate result")
	}

	return &calcModels.CalculateOutput{Body: res}, nil
}

func (h *Handler) History(_ context.Context, _ *struct{}) (*calcModels.HistoryOutput, error) {
	res, err := h.historyService.GetHistory()
	if err != nil {
		return nil, huma.Error500InternalServerError("failed to get calculation history")
	}

	return &calcModels.HistoryOutput{Body: res}, nil
}
