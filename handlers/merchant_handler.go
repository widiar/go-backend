package handlers

import (
	"backendmaw/dto"
	"backendmaw/services"
	"net/http"

	"github.com/labstack/echo/v5"
)

type MerchantHandler struct {
	service *services.MerchantService
}

func NewMerchantHandler(service *services.MerchantService) *MerchantHandler {
	return &MerchantHandler{service: service}
}

func (h *MerchantHandler) List(c *echo.Context) error {
	c.Logger().Info("[START] ListMerchantHandler")
	response, err := h.service.List()
	c.Logger().Info("[END] ListMerchantHandler", "error", err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response)
	}
	return c.JSON(http.StatusOK, response)
}

func (h *MerchantHandler) Create(c *echo.Context) error {
	c.Logger().Info("[START] CreateMerchantHandler")
	var request dto.MerchantRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Input")
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	response, err := h.service.Create(&request)
	c.Logger().Info("[END] CreateMerchantHandler", "error", err)
	return c.JSON(response.Status, response)
}

func (h *MerchantHandler) Update(c *echo.Context) error {
	c.Logger().Info("[START] UpdateMerchantHandler")
	id := c.Param("id")
	var request dto.MerchantRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Input")
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	response, err := h.service.Update(id, &request)
	c.Logger().Info("[END] UpdateMerchantHandler", "error", err)
	return c.JSON(response.Status, response)
}

func (h *MerchantHandler) Delete(c *echo.Context) error {
	c.Logger().Info("[START] DeleteMerchantHandler")
	id := c.Param("id")
	response, err := h.service.Delete(id)
	c.Logger().Info("[END] DeleteMerchantHandler", "error", err)
	return c.JSON(response.Status, response)
}

func (h *MerchantHandler) MerchantFeature(c *echo.Context) error {
	c.Logger().Info("[START] MerchantFeatureHandler")
	var request dto.MerchantFeatureRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Input")
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	response, err := h.service.RelateFeature(&request)
	c.Logger().Info("[END] MerchantFeatureHandler", "error", err)
	return c.JSON(response.Status, response)
}

func (h *MerchantHandler) ListMerchantFeature(c *echo.Context) error {
	c.Logger().Info("[START] ListMerchantFeatureHandler")
	response, err := h.service.ListMerchantFeature()
	c.Logger().Info("[END] ListMerchantFeatureHandler", "error", err)
	return c.JSON(http.StatusOK, response)
}
