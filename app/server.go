package app

import (
	"fmt"
	"golang-echo-wasm-example/app/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunServer(port int) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middlewares.HealthCheckMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"title": "main",
		})
	})

	e.GET("/api/v1/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil && err != http.ErrServerClosed {
		e.Logger.Error(err)
		e.Logger.Panic("shuttig down the server")
	}
}
