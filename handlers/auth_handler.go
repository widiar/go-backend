package handlers

import (
	"backendmaw/dto"
	"backendmaw/services"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

func RegisterHandler(c *echo.Context) error {
	c.Logger().Info("[START] Register")
	register := new(dto.RegisterRequest)
	if err := c.Bind(register); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Input")
	}
	if err := c.Validate(register); err != nil {
		return err
	}
	resp, err := services.RegisterService(register)
	c.Logger().Info("[END] Register", "error", err)
	if err != nil {
		return c.JSON(resp.Status, resp)
	}
	return c.JSON(http.StatusCreated, resp)
}

func LoginHandler(c *echo.Context) error {
	c.Logger().Info("[START] Login")
	login := new(dto.LoginRequest)
	if err := c.Bind(login); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Input")
	}
	if err := c.Validate(login); err != nil {
		return err
	}
	resp, err := services.LoginService(login)
	c.Logger().Info("[END] Login", "error", err)
	if err != nil {
		return c.JSON(resp.Status, resp)
	}
	cookie := &http.Cookie{
		Name:     "token",
		Value:    resp.Payload.(*dto.LoginResponse).Token,
		Expires:  time.Now().Add(1 * time.Hour),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, resp)
}
