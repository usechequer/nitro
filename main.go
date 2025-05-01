package main

import (
	"log"
	"nitro/models"
	"nitro/utilities"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("There was a problem loading the environment variables")
	}

	database := utilities.GetDatabaseObject()

	database.AutoMigrate(&models.Project{}, &models.NotificationConfig{}, &models.StatusPage{})

}
