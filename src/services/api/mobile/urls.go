package mobile

import (
	"app/src/services/api/mobile/calculator"

	"github.com/danielgtaylor/huma/v2"
)

func InitMobile(api huma.API, dependencies *Dependencies) {
	calculator.InitCalculator(api, dependencies.Calculator)
}
