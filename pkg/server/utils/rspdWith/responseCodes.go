package rspdWith

import (
	"github.com/labstack/echo"
	"net/http"
)

type (
	HTTPResponse struct {
		Error   bool        `json:"error"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

func HTTPSuccess(c echo.Context, m string, d interface{}) error {
	return c.JSON(http.StatusOK, HTTPResponse{
		Error:   false,
		Message: m,
		Data:    d,
	})
}