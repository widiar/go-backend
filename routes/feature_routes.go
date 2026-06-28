package routes

import (
	"backendmaw/handlers"

	"github.com/labstack/echo/v5"
)

func FeatureRoutes(g *echo.Group, handlers *handlers.FeatureHandler) {
	g.GET("/features", handlers.List)
	g.POST("/features", handlers.Create)
	g.PUT("/features/:id", handlers.Update)
	g.DELETE("/features/:id", handlers.Delete)
}
