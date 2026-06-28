package handlers

import (
	"backendmaw/services"
	"net/http"

	"github.com/labstack/echo/v5"
)

type BannerHandler struct {
	service *services.BannerService
}

func NewBannerHandler(service *services.BannerService) *BannerHandler {
	return &BannerHandler{service}
}

func (h *BannerHandler) ListBanner(c *echo.Context) error {
	c.Logger().Info("[END] ListBannerHandler")
	resp, err := h.service.ListBanner()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
