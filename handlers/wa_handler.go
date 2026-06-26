package handlers

import (
	"backendmaw/services"
	"net/http"

	"github.com/labstack/echo/v5"
)

func LoginWaHandler(c *echo.Context) error {
	c.Logger().Info("[START] Login WA")
	resp, err := services.LoginWa()
	c.Logger().Info("[END] Login WA", "error", err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
