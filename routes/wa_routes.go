package routes

import (
	"backendmaw/handlers"

	"github.com/labstack/echo/v5"
)

func WaRoutes(g *echo.Group) {
	g.GET("/wa/login", handlers.LoginWaHandler)
}
