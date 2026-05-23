package mobile

import (
	"app/src/services/api/mobile/calculator"
	"net/http"
)

func InitMobile(mux *http.ServeMux) {
	calculatorMux := http.NewServeMux()

	calculator.InitCalculator(calculatorMux)

	mux.Handle("/mobile/", http.StripPrefix("/mobile", calculatorMux))
}
