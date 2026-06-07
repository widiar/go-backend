package middlewares

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func CorrelationLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			reqId := uuid.NewString()
			newLogger := c.Logger().With("correlation_id", reqId)
			c.SetLogger(newLogger)
			return next(c)
		}
	}
}
