package handlers

import (
	"backendmaw/dto"
	"backendmaw/services"
	"net/http"

	"github.com/labstack/echo/v5"
)

type WaHandler struct {
	service *services.WaService
}

func NewWaHandler(service *services.WaService) *WaHandler {
	return &WaHandler{service: service}
}

func (h *WaHandler) Login(c *echo.Context) error {
	c.Logger().Info("[START] Login WA")
	resp, err := h.service.Login()
	c.Logger().Info("[END] Login WA", "error", err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp)
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *WaHandler) SendMessage(c *echo.Context) error {
	c.Logger().Info("[START] SendMessage WA")
	var request dto.SendWaMessageRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	resp, err := h.service.SendMessage(&request)
	c.Logger().Info("[END] SendMessage WA", "error", err)
	return c.JSON(http.StatusOK, resp)
}
