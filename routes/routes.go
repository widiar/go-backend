package routes

import (
	"backendmaw/middlewares"

	"github.com/labstack/echo/v5"
)

func Routes(e *echo.Echo) {
	authGroup := e.Group("/api/auth")
	AuthRoutes(authGroup)

	protectedGroup := e.Group("/api")
	protectedGroup.Use(middlewares.JwtMiddleware)
	BannerRoute(protectedGroup)
	MerchantRoutes(protectedGroup)
	FeatureRoutes(protectedGroup)
	WaRoutes(protectedGroup)
}
