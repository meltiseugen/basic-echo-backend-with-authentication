package routes

import (
	"RevelTest/pkg/server/handler"
	"fmt"
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	e.Static("/static", "pkg/assets")
	e.File("/favicon.ico", "pkg/images/favicon.ico")

	e.GET(GetRoute(METRICS), echo.WrapHandler(promhttp.Handler()))

	e.GET(GetRoute(INDEX), handler.Index)
	e.GET(GetRoute(MESSAGE, V1), handler.UserMessage)
}
