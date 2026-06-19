package routes

import (
	"backendmaw/handlers"

	"github.com/labstack/echo/v5"
)

func MerchantRoutes(g *echo.Group) {
	g.GET("/merchant", handlers.ListMerchantHandler)
	g.POST("/merchant", handlers.CreateMerchantHandler)
	g.GET("/merchant/feature", handlers.ListMerchantFeatureHandler)
	g.POST("/merchant/feature", handlers.MerchantFeatureHandler)
	g.PUT("/merchant/:id", handlers.UpdateMerchantHandler)
	g.DELETE("/merchant/:id", handlers.DeleteMerchantHandler)
}
