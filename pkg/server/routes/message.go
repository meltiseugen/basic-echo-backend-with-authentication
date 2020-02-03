package routes

import (
	"RevelTest/pkg/server/handler"
	"github.com/labstack/echo"
)

func InitMessage(e *echo.Echo) {
	e.GET(GetRoute(MESSAGE, V1), handler.UserMessage)
}
