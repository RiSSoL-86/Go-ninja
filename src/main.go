package main

import (
	"app/src/services/api"
	"log"
	"net/http"
)

func main() {
	apiMux := http.NewServeMux()

	api.InitAPI(apiMux)
	log.Println("Server starting on :8080")

	if err := http.ListenAndServe(":8080", apiMux); err != nil {
		log.Fatal(err)
	}
}
