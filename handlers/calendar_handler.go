package handlers

import (
	"backendmaw/dto"
	"backendmaw/services"
	"net/http"

	"github.com/labstack/echo/v5"
)

type CalendarHandler struct {
	service *services.CalendarService
}

func NewCalendarHandler(service *services.CalendarService) *CalendarHandler {
	return &CalendarHandler{service: service}
}

func (h *CalendarHandler) ListEventHolidayWfh(c *echo.Context) error {
	var request dto.WfhRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	response, err := h.service.EventWfh()
	c.Logger().Info("[END] Service Calendar", "error", err)
	return c.JSON(response.Status, response)
}
