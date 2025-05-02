package validators

import (
	"errors"
	"fmt"
	"net/http"
	"nitro/controllers"
	"nitro/dto"
	"nitro/models"
	"nitro/utilities"
	"strings"

	"github.com/labstack/echo"
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
