package calculator

import (
	"time"

	"github.com/lib/pq"
)

type CalculationRecord struct {
	ID        int64           `gorm:"primaryKey"`
	Args      pq.Float64Array `gorm:"type:float8[]"`
	Operator  string          `gorm:"type:text"`
	Result    float64         `gorm:"type:float8"`
	Note      string          `gorm:"type:text"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`
}

func (CalculationRecord) TableName() string {
	return "calculations"
}
