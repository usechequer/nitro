package main

import (
	"log"
	"nitro/middleware"
	"nitro/models"
	"nitro/validators"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/usechequer/utilities"
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
	app.Use(utilities.AuthMiddleware)
	app.Use(middleware.TokenMiddleware)

	app.POST("/projects", validators.CreateProjectValidator)
	app.PUT("/projects/:uuid", validators.UpdateProjectValidator)

	app.POST("/projects/:project_uuid/notification-configs", validators.ValidateCreateNotificationConfigDto)
	app.PUT("/projects/:project_uuid/notification-configs/:uuid", validators.ValidateUpdateNotificationConfig)

	app.POST("/projects/:project_uuid/status-pages", validators.ValidateCreateStatusPage)
	app.PUT("/projects/:project_uuid/status-pages/:uuid", validators.ValidateUpdateStatusPage)

	app.Logger.Fatal(app.Start(":8000"))
}
