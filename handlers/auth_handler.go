package handlers

import (
	"backendmaw/dto"
	"backendmaw/services"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Register(c *echo.Context) error {
	c.Logger().Info("[START] Register")
	register := new(dto.RegisterRequest)
	if err := c.Bind(register); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Input")
	}
	if err := c.Validate(register); err != nil {
		return err
	}
	resp, err := h.AuthService.Register(register)
	c.Logger().Info("[END] Register", "error", err)
	if err != nil {
		return c.JSON(resp.Status, resp)
	}
	return c.JSON(http.StatusCreated, resp)
}

func (h *AuthHandler) Login(c *echo.Context) error {
	c.Logger().Info("[START] Login")
	login := new(dto.LoginRequest)
	if err := c.Bind(login); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Input")
	}
	if err := c.Validate(login); err != nil {
		return err
	}
	resp, err := h.AuthService.Login(login)
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

func (h *AuthHandler) Me(c *echo.Context) error {
	c.Logger().Info("[START] Me")
	resp, err := h.AuthService.Me(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp)
	}
	c.Logger().Info("[END] Me", "error", err)
	return c.JSON(http.StatusOK, resp)
}

func (h *AuthHandler) Logout(c *echo.Context) error {
	c.Logger().Info("[START] Logout")
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-1 * time.Hour),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, "Success")
}
