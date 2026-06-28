package routes

import (
	"backendmaw/handlers"

	"github.com/labstack/echo/v5"
)

func WaRoutes(g *echo.Group, waHandler *handlers.WaHandler) {
	g.GET("/wa/login", waHandler.Login)
	g.POST("/wa/send", waHandler.SendMessage)
}
