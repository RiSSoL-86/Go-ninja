package main

import (
	"app/src/app_settings"
	"app/src/core/db"
	"app/src/services/api"
	"log"
	"net/http"
)

func main() {
	config, err := app_settings.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	dsn := config.Database.GetDSN()

	// Connect to DB using GORM
	dbConn, err := db.ConnectGORM(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	apiMux := http.NewServeMux()
	dependencies := api.NewDependencies(dbConn)

	api.InitAPI(apiMux, dependencies, config.API.HumaConfig())
	log.Println("Server starting on :8080")

	if err := http.ListenAndServe(":8080", apiMux); err != nil {
		log.Fatal(err)
	}
}
