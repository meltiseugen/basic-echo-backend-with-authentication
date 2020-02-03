package handler

import (
	"RevelTest/pkg/server/utils"
	"RevelTest/pkg/server/utils/rspdWith"
	"github.com/labstack/echo"
)

func Index(c echo.Context) (err error) {
	defer utils.Log("handling Index...")
	return rspdWith.HTTPSuccess(c, "welcome!", nil)
}