package calculator

import (
	calcModels "app/src/services/api/mobile/calculator/models"
)

type MemoryRepository struct{}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{}
}

func (r *MemoryRepository) SaveCalculation(req calcModels.CalculateRequest, res calcModels.CalculateResponse) error {
	// Имитация сохранения в БД
	return nil
}

func (r *MemoryRepository) GetHistory() ([]calcModels.CalculateRequest, error) {
	// Имитация получения данных из БД
	return []calcModels.CalculateRequest{}, nil
}
