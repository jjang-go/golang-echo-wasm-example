package handlers

import (
	"github.com/labstack/echo/v4"
)

func MainHandler(ctx echo.Context) error {
	data := map[string]interface{}{
		"Title":   "Welcome to Echo",
		"Message": "Hello, Echo with Templates!",
	}
	return ctx.Render(200, "index.html", data)
}
