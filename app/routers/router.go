package routers

import (
	"golang-echo-wasm-example/app/handlers"
	v1 "golang-echo-wasm-example/app/handlers/api/v1"
	v2 "golang-echo-wasm-example/app/handlers/api/v2"

	"github.com/labstack/echo/v4"
)

func MainRoute(g *echo.Group) {
	g.GET("/", handlers.MainHandler)
}

func ApiV1Route(g *echo.Group) {
	g.GET("/hello", v1.ContentV1Handler)
}

func ApiV2Route(g *echo.Group) {
	g.GET("/hello", v2.ContentV2Handler)
}
