package routes

import (
	"backendmaw/handlers"

	"github.com/labstack/echo/v5"
)

func AuthRoutes(g *echo.Group) {
	g.POST("/auth/register", handlers.RegisterHandler)
	g.POST("/auth/login", handlers.LoginHandler)
}
