package calculator

import (
	"app/src/services/api/common/handler"
	calcModels "app/src/services/api/mobile/calculator/models"
	"encoding/json"
	"net/http"
)

type ICalculateService interface {
	Execute(req calcModels.CalculateRequest) (calcModels.CalculateResponse, error)
}

type IHistoryService interface {
	GetHistory() ([]calcModels.CalculateRequest, error)
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

func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {
	var req calcModels.CalculateRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	res, err := h.calculateService.Execute(req)
	if err != nil {
		handler.Response(w, calcModels.CalculateResponse{
			Error: err.Error(),
		}, http.StatusUnprocessableEntity)
		return
	}

	handler.Response(w, res, http.StatusOK)
}

func (h *Handler) History(w http.ResponseWriter, r *http.Request) {
	res, err := h.historyService.GetHistory()
	if err != nil {
		handler.Response(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	handler.Response(w, res, http.StatusOK)
}
