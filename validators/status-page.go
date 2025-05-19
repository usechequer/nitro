package validators

import (
	"errors"
	"fmt"
	"net/http"
	"nitro/controllers"
	"nitro/dto"
	"nitro/models"

	"github.com/labstack/echo/v4"
	"github.com/usechequer/utilities"
	"gorm.io/gorm"
)

func ValidateCreateStatusPage(context echo.Context) error {
	createStatusPageDto := new(dto.CreateStatusPageDto)

	if err := context.Bind(createStatusPageDto); err != nil {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusBadRequest, Error: "MALFORMED_REQUEST", Message: err.Error()})
	}

	if err := context.Validate(createStatusPageDto); err != nil {
		return err
	}

	var project models.Project
	database := utilities.GetDatabaseObject()

	projectResult := database.Where("uuid = ?", createStatusPageDto.ProjectUuid.String()).First(&project)

	if errors.Is(projectResult.Error, gorm.ErrRecordNotFound) {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusNotFound, Error: "PROJECT_002", Message: fmt.Sprintf("Project with uuid %s does not exist", createStatusPageDto.ProjectUuid)})
	}

	var statusPage models.StatusPage
	statusPageResult := database.Where("project_id = ?", project.ID).First(&statusPage)

	if !errors.Is(statusPageResult.Error, gorm.ErrRecordNotFound) {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusBadRequest, Error: "STATUS_PAGE_001", Message: fmt.Sprintf("Project with uuid %s already has an existing status page", createStatusPageDto.ProjectUuid)})
	}

	context.Set("project", project)
	context.Set("createStatusPageDto", createStatusPageDto)

	return controllers.CreateStatusPage(context)
}

func ValidateUpdateStatusPage(context echo.Context) error {
	updateStatusPageDto := new(dto.UpdateStatusPageDto)

	if err := context.Bind(updateStatusPageDto); err != nil {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusBadRequest, Error: "MALFORMED_REQUEST", Message: err.Error()})
	}

	if err := context.Validate(updateStatusPageDto); err != nil {
		return err
	}

	var statusPage models.StatusPage
	database := utilities.GetDatabaseObject()

	result := database.Preload("Project").Where("uuid = ?", updateStatusPageDto.Uuid).First(&statusPage)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusNotFound, Error: "STATUS_PAGE_002", Message: fmt.Sprintf("Status page with uuid %s does not exist", updateStatusPageDto.Uuid)})
	}

	if statusPage.Project.Uuid.String() != updateStatusPageDto.ProjectUuid.String() {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusNotFound, Error: "STATUS_PAGE_003", Message: fmt.Sprintf("Status page with uuid %s is not tied to project with uuid %s", updateStatusPageDto.Uuid, updateStatusPageDto.ProjectUuid)})
	}

	context.Set("statusPage", statusPage)
	context.Set("updateStatusPageDto", updateStatusPageDto)

	return controllers.UpdateStatusPage(context)
}
