package handlers

import (
	"backendmaw/dto"
	"backendmaw/services"
	"net/http"

	"github.com/labstack/echo/v5"
)

func ListFeatureHandler(c *echo.Context) error {
	c.Logger().Info("[START] ListMerchantHandler")
	response, err := services.ListFeatureService()
	c.Logger().Info("[END] ListMerchantHandler", "error", err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response)
	}
	return c.JSON(http.StatusOK, response)
}

func CreateFeatureHandler(c *echo.Context) error {
	c.Logger().Info("[START] CreateFeatureHandler")
	var request dto.FeatureRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	response, err := services.CreateFeatureService(&request)
	c.Logger().Info("[END] CreateFeatureHandler", "error", err)
	return c.JSON(response.Status, response)
}

func UpdateFeatureHandler(c *echo.Context) error {
	c.Logger().Info("[START] UpdateFeatureHandler")
	var request dto.FeatureRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	id := c.Param("id")
	response, err := services.UpdateFeatureService(id, &request)
	c.Logger().Info("[END] UpdateFeatureHandler", "error", err)
	return c.JSON(response.Status, response)
}

func DeleteFeatureHandler(c *echo.Context) error {
	c.Logger().Info("[START] DeleteFeatureHandler")
	id := c.Param("id")
	response, err := services.DeleteFeatureService(id)
	c.Logger().Info("[END] DeleteFeatureHandler", "error", err)
	return c.JSON(response.Status, response)
}
