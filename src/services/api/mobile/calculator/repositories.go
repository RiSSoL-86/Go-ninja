package calculator

import (
	"app/src/core/models/calculator"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SaveCalculation(record *calculator.CalculationRecord) error {
	return r.db.Create(record).Error
}

func (r *Repository) GetHistory() ([]calculator.CalculationRecord, error) {
	var history []calculator.CalculationRecord
	err := r.db.Order("created_at DESC").Find(&history).Error
	return history, err
}
