package controllers

import (
	"net/http"
	"nitro/dto"
	"nitro/models"

	"github.com/labstack/echo/v4"
	"github.com/usechequer/utilities"
)

func CreateNotificationConfig(context echo.Context) error {
	project := context.Get("project").(models.Project)
	createNotificationConfigDto := context.Get("CreateNotificationConfigDto").(*dto.CreateNotificationConfigDto)

	notificationConfig := models.NotificationConfig{Config: createNotificationConfigDto.Config, ProjectID: project.ID}

	database := utilities.GetDatabaseObject()
	result := database.Create(&notificationConfig)

	if result.Error != nil {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusInternalServerError, Error: "INTERNAL_SERVER_ERROR", Message: result.Error.Error()})
	}

	return context.JSON(http.StatusCreated, notificationConfig)
}

func UpdateNotificationConfig(context echo.Context) error {
	notificationConfig := context.Get("notificationConfig").(models.NotificationConfig)
	updateNotificationConfigDto := context.Get("updateNotificationConfigDto").(*dto.UpdateNotificationConfigDto)

	notificationConfig.Config = updateNotificationConfigDto.Config
	database := utilities.GetDatabaseObject()

	result := database.Save(&notificationConfig)

	if result.Error != nil {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusInternalServerError, Error: "INTERNAL_SERVER_ERROR", Message: result.Error.Error()})
	}

	return context.JSON(200, notificationConfig)
}
