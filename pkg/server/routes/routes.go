package routes

import (
	"fmt"
	"github.com/labstack/echo"
)

const (
	V1 string = "v1"
	V2        = "v2"
)

type Route string

const (
	INDEX   Route = "/"
	METRICS       = INDEX + "metrics"
	ROOT          = INDEX + "%s/"

	MESSAGE = ROOT + "message"
)

func GetRoute(route Route, params ...string) string {
	if len(params) == 0 {
		return fmt.Sprintf(string(route))
	}

	return fmt.Sprintf(string(route), params[0])
}

func Init(e *echo.Echo) {
	InitSystem(e)
	InitIndex(e)
	InitMessage(e)
	InitInstrumenting(e)
}
