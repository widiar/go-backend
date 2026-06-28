package routes

import (
	"backendmaw/handlers"

	"github.com/labstack/echo/v5"
)

func MerchantRoutes(g *echo.Group, handlers *handlers.MerchantHandler) {
	g.GET("/merchant", handlers.List)
	g.POST("/merchant", handlers.Create)
	g.GET("/merchant/feature", handlers.ListMerchantFeature)
	g.POST("/merchant/feature", handlers.MerchantFeature)
	g.PUT("/merchant/:id", handlers.Update)
	g.DELETE("/merchant/:id", handlers.Delete)
}
