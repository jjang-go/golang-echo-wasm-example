package routers

import (
	"golang-echo-wasm-example/app/handlers"

	"github.com/labstack/echo/v4"
)

func MainRoute(g *echo.Group) {
	g.GET("/", handlers.MainHandler)
	g.GET("/health", handlers.MainHealthCheck)
}
