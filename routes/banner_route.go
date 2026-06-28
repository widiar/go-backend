package routes

import (
	"backendmaw/handlers"

	"github.com/labstack/echo/v5"
)

func BannerRoute(g *echo.Group, handler *handlers.BannerHandler) {
	g.GET("/banner", handler.ListBanner)
}
