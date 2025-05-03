package controllers

import (
	"net/http"
	"nitro/dto"
	"nitro/models"
	"nitro/utilities"

	"github.com/labstack/echo"
)

func CreateProject(context echo.Context) error {
	createProjectDto := context.Get("createProjectDto").(*dto.CreateProjectDto)

	var descriptionPointer *string

	if len(createProjectDto.Description) > 0 {
		descriptionPointer = &createProjectDto.Description
	}

	project := models.Project{UserUuid: createProjectDto.UserUuid, Name: createProjectDto.Name, Url: createProjectDto.Url, Description: descriptionPointer}
	database := utilities.GetDatabaseObject()
	result := database.Create(&project)

	if result.Error != nil {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusInternalServerError, Error: "INTERNAL_SERVER_ERROR", Message: result.Error.Error()})
	}

	return context.JSON(http.StatusCreated, project)
}

func UpdateProject(context echo.Context) error {
	// project := context.Get("project").(*models.Project)
	// updateProjectDto := context.Get("updateProjectDto").(*dto.UpdateProjectDto)

	return context.JSON(123, map[string]string{})
}
