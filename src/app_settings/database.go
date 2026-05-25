package app_settings

import (
	"errors"
	"os"
	"strings"
)

type DatabaseSettings struct {
	DSN string
}

func NewDatabaseSettings() (*DatabaseSettings, error) {
	dsn := strings.TrimSpace(os.Getenv("DATABASE_URL"))
	if dsn == "" {
		return nil, errors.New("DATABASE_URL is not set")
	}

	return &DatabaseSettings{DSN: dsn}, nil
}

func (s *DatabaseSettings) GetDSN() string {
	return s.DSN
}
