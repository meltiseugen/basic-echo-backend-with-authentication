package server

import (
	"RevelTest/pkg/server/routes"
	"RevelTest/pkg/server/utils"
	"RevelTest/pkg/server/utils/auth"
	"context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SetUpRoutes(e *echo.Echo) {
	defer utils.Log("Set up the routes")


}

func SetUpLogger(e *echo.Echo) {
	defer utils.Log("Set up the logger")

	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
}

func SetDefaultFlags(e *echo.Echo) {
	e.HideBanner = true
	e.Pre(middleware.RemoveTrailingSlash())

}

func SetUpJWTAuth(e * echo.Echo) {
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:        &auth.JWTClaims{},
		SigningKey:    []byte("secret"),
		SigningMethod: "HS512",
		Skipper: func(c echo.Context) bool {
			return utils.Contains(auth.Excluded, c.Path())
		},
	}))
}

func Start() {
	e := echo.New()

	SetDefaultFlags(e)
	SetUpLogger(e)
	SetUpJWTAuth(e)
	SetUpRoutes(e)

	// Start the server
	go func() {
		if err := e.Start(":1323"); err != nil {
			e.Logger.Info(err)
		}
	}()

	// Graceful shutdown of the server with a timeout
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-quit
	e.Logger.Info("gracefully shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}