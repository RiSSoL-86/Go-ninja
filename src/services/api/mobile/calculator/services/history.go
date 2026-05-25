package services

import (
	"app/src/core/models/calculator"
	"app/src/services/api/mobile/calculator/models"
)

type IHistoryRepository interface {
	GetHistory() ([]calculator.CalculationRecord, error)
}

type HistoryService struct {
	repository IHistoryRepository
}

func NewHistoryService(repository IHistoryRepository) *HistoryService {
	return &HistoryService{repository: repository}
}

func (s *HistoryService) GetHistory() ([]models.HistoryItem, error) {
	records, err := s.repository.GetHistory()
	if err != nil {
		return nil, err
	}

	history := make([]models.HistoryItem, 0, len(records))
	for _, record := range records {
		history = append(history, models.HistoryItem{
			ID:        record.ID,
			Args:      append([]float64(nil), record.Args...),
			Operator:  record.Operator,
			Result:    record.Result,
			Note:      record.Note,
			CreatedAt: record.CreatedAt,
		})
	}

	return history, nil
}
