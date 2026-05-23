package calculator

import (
	"net/http"
)

func InitCalculator(mux *http.ServeMux) {
	calculator := BuildCalculator()

	mux.HandleFunc("POST /calculate", calculator.Calculate)
	mux.HandleFunc("GET /history", calculator.History)
}
