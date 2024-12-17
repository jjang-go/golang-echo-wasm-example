package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ContentV1Handler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "api v1 OK")
}
