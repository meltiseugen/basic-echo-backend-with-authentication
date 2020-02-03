package routes

import (
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitInstrumenting(e *echo.Echo) {
	e.GET(GetRoute(METRICS), echo.WrapHandler(promhttp.Handler()))
}
