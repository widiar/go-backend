package routes

import (
	"github.com/labstack/echo/v5"
)

func Routes(e *echo.Echo) {
	api := e.Group("/api")
	AuthRoutes(api)
}
