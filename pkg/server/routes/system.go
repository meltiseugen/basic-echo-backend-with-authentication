package routes

import (
	"github.com/labstack/echo"
)

func InitSystem(e *echo.Echo) {
	e.Static("/static", "pkg/assets")
	e.File("/favicon.ico", "pkg/images/favicon.ico")
}
