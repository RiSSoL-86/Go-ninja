package calculator

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func InitCalculator(api huma.API, dependencies *Dependencies) {
	huma.Register(api, huma.Operation{
		OperationID: "calculate",
		Method:      http.MethodPost,
		Path:        "/api/mobile/calculate",
		Summary:     "Calculate and store a result",
	}, dependencies.handler.Calculate)
	huma.Register(api, huma.Operation{
		OperationID: "calculation-history",
		Method:      http.MethodGet,
		Path:        "/api/mobile/history",
		Summary:     "List stored calculations",
	}, dependencies.handler.History)
}
