package main

import (
	"log"
	"nitro/models"
	"nitro/utilities"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("There was a problem loading the environment variables")
	}

	database := utilities.GetDatabaseObject()

	database.AutoMigrate(&models.Project{}, &models.NotificationConfig{}, &models.StatusPage{})

	app := echo.New()
	app.Validator = &utilities.RequestValidator{Validator: validator.New()}

	app.Logger.Fatal(app.Start(":8000"))
}
