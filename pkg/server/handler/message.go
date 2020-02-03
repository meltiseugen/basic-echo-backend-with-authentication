package handler

import (
	"RevelTest/pkg/server/utils/rspdWith"
	"github.com/labstack/echo"
)

func UserMessage(c echo.Context) (err error) {
	return rspdWith.HTTPSuccess(c, "Hello!", nil)
}