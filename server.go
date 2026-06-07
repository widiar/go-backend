package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	// load env
	if err := godotenv.Load(); err != nil {
		e.Logger.Error("Error loading .env file!", "error", err)
		return
	}

	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		e.Logger.Error("Failed to start server", "error", err)
	}
}
