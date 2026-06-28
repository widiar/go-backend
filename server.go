package main

import (
	"backendmaw/config"
	"backendmaw/handlers"
	"backendmaw/middlewares"
	"backendmaw/routes"
	"fmt"
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
	if os.Getenv("DB_HOST") == "" {
		if err := godotenv.Load(); err != nil {
			e.Logger.Error("Error loading .env file!", "error", err)
			panic(err)
		}
		e.Logger.Info("Using environment variable from .env file")
	} else {
		e.Logger.Info("Using environment variable from Docker")
	}

	//load db
	db := config.ConnectDB()
	sqlDb, _ := db.DB()
	defer func() {
		if err := sqlDb.Close(); err != nil {
			fmt.Println("Failed to close DB")
		} else {
			fmt.Println("DB closed safely")
		}
	}()

	waContainer := config.SetupWa(sqlDb)
	handler := handlers.Setup(db, waContainer)
	routes.Setup(e, handler)
	e.Validator = config.NewCustomValidator()
	e.HTTPErrorHandler = config.SetupHttpErrorHandler

	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		e.Logger.Error("Failed to start server", "error", err)
	}
}
