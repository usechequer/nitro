package controllers

import (
	"net/http"
	"nitro/dto"
	"nitro/models"

	"github.com/labstack/echo/v4"
	"github.com/usechequer/utilities"
)

func CreateStatusPage(context echo.Context) error {
	project := context.Get("project").(models.Project)
	createStatusPageDto := context.Get("createStatusPageDto").(*dto.CreateStatusPageDto)

	statusPage := models.StatusPage{Title: createStatusPageDto.Title, Description: createStatusPageDto.Description, Metadata: createStatusPageDto.Metadata, ProjectID: project.ID, Config: createStatusPageDto.Config}

	database := utilities.GetDatabaseObject()
	result := database.Save(&statusPage)

	if result.Error != nil {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusInternalServerError, Error: "INTERNAL_SERVER_ERROR", Message: result.Error.Error()})
	}

	return context.JSON(http.StatusCreated, statusPage)
}

func UpdateStatusPage(context echo.Context) error {
	statusPage := context.Get("statusPage").(models.StatusPage)
	updateStatusPageDto := context.Get("updateStatusPageDto").(*dto.UpdateStatusPageDto)

	updateStringField := func(field *string, param *string) *string {
		if param == nil {
			return field
		}

		return param
	}

	if updateStatusPageDto.Metadata != nil {
		statusPage.Metadata = updateStatusPageDto.Metadata
	}

	if updateStatusPageDto.Config != nil {
		statusPage.Config = *updateStatusPageDto.Config
	}

	statusPage.Title = updateStringField(statusPage.Title, updateStatusPageDto.Title)
	statusPage.Description = updateStringField(statusPage.Description, updateStatusPageDto.Description)

	database := utilities.GetDatabaseObject()
	result := database.Save(&statusPage)

	if result.Error != nil {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusInternalServerError, Error: "INTERNAL_SERVER_ERROR", Message: result.Error.Error()})
	}

	return context.JSON(http.StatusOK, statusPage)
}
