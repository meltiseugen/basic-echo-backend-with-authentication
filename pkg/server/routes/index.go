package routes

import (
	"RevelTest/pkg/server/handler"
	"github.com/labstack/echo"
)

func InitIndex(e *echo.Echo) {
	e.GET(GetRoute(INDEX), handler.Index)
}
