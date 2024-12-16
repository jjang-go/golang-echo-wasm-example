package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func HealthCheckMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasSuffix(c.Request().URL.Path, "/health") {
			return c.JSON(http.StatusOK, map[string]string{
				"status": "UP",
			})
		}
		return next(c)
	}
}
