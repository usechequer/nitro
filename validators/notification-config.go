package validators

import (
	"errors"
	"fmt"
	"net/http"
	"nitro/controllers"
	"nitro/dto"
	"nitro/models"
	"nitro/utilities"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ValidateCreateNotificationConfig(context echo.Context) error {
	createNotificationConfigDto := new(dto.CreateNotificationConfig)

	if err := context.Bind(createNotificationConfigDto); err != nil {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusBadRequest, Error: "MALFORMED_REQUEST", Message: err.Error()})
	}

	if err := context.Validate(createNotificationConfigDto); err != nil {
		return err
	}

	var project models.Project
	database := utilities.GetDatabaseObject()

	projectResult := database.Where("uuid = ?", createNotificationConfigDto.ProjectUuid).First(&project)

	if errors.Is(projectResult.Error, gorm.ErrRecordNotFound) {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusNotFound, Error: "PROJECT_002", Message: fmt.Sprintf("Project with uuid %s does not exist", createNotificationConfigDto.ProjectUuid)})
	}

	var notificationConfig models.NotificationConfig

	notificationResult := database.Where("project_id = ?", project.ID).First(&notificationConfig)

	if !errors.Is(notificationResult.Error, gorm.ErrRecordNotFound) {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusBadRequest, Error: "NOTIFICATION_CONFIG_001", Message: fmt.Sprintf("Project with uuid %s already has an existing notification config", createNotificationConfigDto.ProjectUuid)})
	}

	context.Set("createNotificationConfigDto", createNotificationConfigDto)
	context.Set("project", project)

	return controllers.CreateNotificationConfig(context)
}
