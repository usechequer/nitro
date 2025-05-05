package controllers

import (
	"net/http"
	"nitro/dto"
	"nitro/models"
	"nitro/utilities"

	"github.com/labstack/echo/v4"
)

func CreateNotificationConfig(context echo.Context) error {
	project := context.Get("project").(models.Project)
	createNotificationConfigDto := context.Get("createNotificationConfigDto").(*dto.CreateNotificationConfig)

	notificationConfig := models.NotificationConfig{Config: createNotificationConfigDto.Config, ProjectID: project.ID}

	database := utilities.GetDatabaseObject()
	result := database.Create(&notificationConfig)

	if result.Error != nil {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusInternalServerError, Error: "INTERNAL_SERVER_ERROR", Message: result.Error.Error()})
	}

	return context.JSON(http.StatusCreated, notificationConfig)
}
