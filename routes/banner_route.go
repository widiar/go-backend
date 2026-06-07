package routes

import (
	"backendmaw/handlers"

	"github.com/labstack/echo/v5"
)

func BannerRoute(g *echo.Group) {
	g.GET("/banner", handlers.ListBannerHandler)
}
