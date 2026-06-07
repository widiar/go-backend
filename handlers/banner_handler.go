package handlers

import (
	"backendmaw/services"
	"net/http"

	"github.com/labstack/echo/v5"
)

func ListBannerHandler(c *echo.Context) error {
	c.Logger().Info("[END] ListBannerHandler")
	resp, err := services.ListBannersService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
