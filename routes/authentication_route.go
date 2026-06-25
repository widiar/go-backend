package routes

import (
	"backendmaw/handlers"
	"backendmaw/middlewares"

	"github.com/labstack/echo/v5"
)

func AuthRoutes(g *echo.Group) {
	g.POST("/register", handlers.RegisterHandler)
	g.POST("/login", handlers.LoginHandler)
	g.POST("/logout", handlers.LogoutHandler)
	g.GET("/me", handlers.MeHandler, middlewares.JwtMiddleware)
}
