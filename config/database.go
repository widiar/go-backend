package config

import (
	"backendmaw/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, dbUser, dbPassword, dbName, dbPort)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("ERROR Connect to DB")
	}

	if errMigrate := database.AutoMigrate(
		&models.Users{},
		&models.Merchant{},
		&models.Feature{}); errMigrate != nil {
		fmt.Println("ERROR Auto migrate", errMigrate)
		panic("ERROR AutoMigrate DB")
	}

	DB = database
}
