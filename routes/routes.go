package routes

import (
	"backendmaw/handlers"
	"backendmaw/middlewares"

	"github.com/labstack/echo/v5"
)

func Setup(e *echo.Echo, h *handlers.Handlers) {
	authGroup := e.Group("/api/auth")
	AuthRoutes(authGroup, h.Auth)

	protectedGroup := e.Group("/api")
	protectedGroup.Use(middlewares.JwtMiddleware)
	BannerRoute(protectedGroup, h.Banner)
	FeatureRoutes(protectedGroup, h.Feature)
	MerchantRoutes(protectedGroup, h.Merchant)
	WaRoutes(protectedGroup, h.Wa)
}
