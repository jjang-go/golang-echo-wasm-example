package v2

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ContentV2Handler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "api v2 OK")
}
