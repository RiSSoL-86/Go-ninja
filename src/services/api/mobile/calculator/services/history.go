package services

import (
	"app/src/services/api/mobile/calculator/models"
)

type IHistoryRepository interface {
	GetHistory() ([]models.CalculateRequest, error)
}

type HistoryService struct {
	repository IHistoryRepository
}

func NewHistoryService(repository IHistoryRepository) *HistoryService {
	return &HistoryService{repository: repository}
}

func (s *HistoryService) GetHistory() ([]models.CalculateRequest, error) {
	// В будущем здесь может быть какая-то логика фильтрации или обработки
	return s.repository.GetHistory()
}
