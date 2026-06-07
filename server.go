package main

import (
	"backendmaw/config"
	"backendmaw/middlewares"
	"backendmaw/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := echo.New()
	e.Use(middlewares.CorrelationLogger())
	e.Use(middleware.RequestLogger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
	}))

	// load env
	if err := godotenv.Load(); err != nil {
		e.Logger.Error("Error loading .env file!", "error", err)
		return
	}

	config.ConnectDB()
	routes.Routes(e)
	e.Validator = config.NewCustomValidator()
	e.HTTPErrorHandler = config.SetupHttpErrorHandler

	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		e.Logger.Error("Failed to start server", "error", err)
	}
}
