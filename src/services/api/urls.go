package api

import (
	"app/src/services/api/mobile"
	"net/http"
)

func InitAPI(mux *http.ServeMux) {
	mobileMux := http.NewServeMux()

	mobile.InitMobile(mobileMux)

	mux.Handle("/api/", http.StripPrefix("/api", mobileMux))
}
