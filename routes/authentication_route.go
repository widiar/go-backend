package routes

import (
	"backendmaw/handlers"
	"backendmaw/middlewares"

	"github.com/labstack/echo/v5"
)

func AuthRoutes(g *echo.Group, handler *handlers.AuthHandler) {
	g.POST("/register", handler.Register)
	g.POST("/login", handler.Login)
	g.POST("/logout", handler.Logout)
	g.GET("/me", handler.Me, middlewares.JwtMiddleware)
}
