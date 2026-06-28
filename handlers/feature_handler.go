package handlers

import (
	"backendmaw/dto"
	"backendmaw/services"
	"net/http"

	"github.com/labstack/echo/v5"
)

type FeatureHandler struct {
	service *services.FeatureService
}

func NewFeatureHandler(service *services.FeatureService) *FeatureHandler {
	return &FeatureHandler{service: service}
}

func (h *FeatureHandler) List(c *echo.Context) error {
	c.Logger().Info("[START] ListMerchantHandler")
	response, err := h.service.List()
	c.Logger().Info("[END] ListMerchantHandler", "error", err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response)
	}
	return c.JSON(http.StatusOK, response)
}

func (h *FeatureHandler) Create(c *echo.Context) error {
	c.Logger().Info("[START] CreateFeatureHandler")
	var request dto.FeatureRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	response, err := h.service.Create(&request)
	c.Logger().Info("[END] CreateFeatureHandler", "error", err)
	return c.JSON(response.Status, response)
}

func (h *FeatureHandler) Update(c *echo.Context) error {
	c.Logger().Info("[START] UpdateFeatureHandler")
	var request dto.FeatureRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	id := c.Param("id")
	response, err := h.service.Update(id, &request)
	c.Logger().Info("[END] UpdateFeatureHandler", "error", err)
	return c.JSON(response.Status, response)
}

func (h *FeatureHandler) Delete(c *echo.Context) error {
	c.Logger().Info("[START] DeleteFeatureHandler")
	id := c.Param("id")
	response, err := h.service.Delete(id)
	c.Logger().Info("[END] DeleteFeatureHandler", "error", err)
	return c.JSON(response.Status, response)
}
