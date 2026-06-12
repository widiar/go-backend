package routes

import (
	"backendmaw/handlers"

	"github.com/labstack/echo/v5"
)

func FeatureRoutes(g *echo.Group) {
	g.GET("/features", handlers.ListFeatureHandler)
	g.POST("/features", handlers.CreateFeatureHandler)
	g.PUT("/features/:id", handlers.UpdateFeatureHandler)
	g.DELETE("/features/:id", handlers.DeleteFeatureHandler)
}
