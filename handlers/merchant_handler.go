package handlers

import (
	"backendmaw/dto"
	"backendmaw/services"
	"net/http"

	"github.com/labstack/echo/v5"
)

func ListMerchantHandler(c *echo.Context) error {
	c.Logger().Info("[START] ListMerchantHandler")
	response, err := services.ListMerchantService()
	c.Logger().Info("[END] ListMerchantHandler", "error", err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response)
	}
	return c.JSON(http.StatusOK, response)
}

func CreateMerchantHandler(c *echo.Context) error {
	c.Logger().Info("[START] CreateMerchantHandler")
	var request dto.MerchantRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Input")
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	response, err := services.CreateMerchantService(&request)
	c.Logger().Info("[END] CreateMerchantHandler", "error", err)
	return c.JSON(response.Status, response)
}

func UpdateMerchantHandler(c *echo.Context) error {
	c.Logger().Info("[START] UpdateMerchantHandler")
	id := c.Param("id")
	var request dto.MerchantRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Input")
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	response, err := services.UpdateMerchantService(id, &request)
	c.Logger().Info("[END] UpdateMerchantHandler", "error", err)
	return c.JSON(response.Status, response)
}

func DeleteMerchantHandler(c *echo.Context) error {
	c.Logger().Info("[START] DeleteMerchantHandler")
	id := c.Param("id")
	response, err := services.DeleteMerchantService(id)
	c.Logger().Info("[END] DeleteMerchantHandler", "error", err)
	return c.JSON(response.Status, response)
}
