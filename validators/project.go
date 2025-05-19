package validators

import (
	"errors"
	"fmt"
	"net/http"
	"nitro/controllers"
	"nitro/dto"
	"nitro/models"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/usechequer/utilities"
	"gorm.io/gorm"
)

func CreateProjectValidator(context echo.Context) error {
	createProjectDto := new(dto.CreateProjectDto)

	if err := context.Bind(createProjectDto); err != nil {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusBadRequest, Error: "MALFORMED_REQUEST", Message: err.Error()})
	}

	if err := context.Validate(createProjectDto); err != nil {
		return err
	}

	createProjectDto.Name = strings.ToLower(createProjectDto.Name)
	database := utilities.GetDatabaseObject()

	var project models.Project
	result := database.Where("user_uuid = ?", createProjectDto.UserUuid).Where("name = ?", createProjectDto.Name).First(&project)

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusBadRequest, Error: "PROJECT_001", Message: fmt.Sprintf("Project with name %s already exists for the specified user", createProjectDto.Name)})
	}

	context.Set("createProjectDto", createProjectDto)

	return controllers.CreateProject(context)
}

func UpdateProjectValidator(context echo.Context) error {
	updateProjectDto := new(dto.UpdateProjectDto)

	if err := context.Bind(updateProjectDto); err != nil {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusBadRequest, Error: "MALFORMED_REQUEST", Message: err.Error()})
	}

	if err := context.Validate(updateProjectDto); err != nil {
		return err
	}

	var project models.Project
	database := utilities.GetDatabaseObject()

	result := database.Where("uuid = ?", updateProjectDto.Uuid).First(&project)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusNotFound, Error: "PROJECT_002", Message: fmt.Sprintf("Project with uuid %s does not exist", updateProjectDto.Uuid)})
	}

	if len(updateProjectDto.Name) > 0 {
		updateProjectDto.Name = strings.ToLower(updateProjectDto.Name)
	}

	context.Set("project", project)
	context.Set("updateProjectDto", updateProjectDto)

	return controllers.UpdateProject(context)
}
