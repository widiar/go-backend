package routes

import (
	"backendmaw/handlers"

	"github.com/labstack/echo/v5"
)

func CalendarRoutes(g *echo.Group, handler *handlers.CalendarHandler) {
	g.GET("/wfh", handler.ListEventHolidayWfh)
	g.POST("/config", handler.ConfigCalendar)
}
