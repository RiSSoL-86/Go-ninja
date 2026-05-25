package app_settings

import "github.com/joho/godotenv"

type Config struct {
	Database *DatabaseSettings
	API      *APISettings
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load("src/.env")

	database, err := NewDatabaseSettings()
	if err != nil {
		return nil, err
	}

	return &Config{
		Database: database,
		API:      NewAPISettings(),
	}, nil
}
