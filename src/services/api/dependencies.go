package api

import (
	"app/src/services/api/mobile"

	"gorm.io/gorm"
)

type Dependencies struct {
	Mobile *mobile.Dependencies
}

func NewDependencies(db *gorm.DB) *Dependencies {
	return &Dependencies{
		Mobile: mobile.NewDependencies(db),
	}
}
