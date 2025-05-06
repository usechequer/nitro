package controllers

import (
	"net/http"
	"nitro/dto"
	"nitro/models"
	"nitro/utilities"

	"github.com/labstack/echo/v4"
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
	// statusPage := context.Get("statusPage").(models.StatusPage)
	// updateStatusPageDto := context.Get("updateStatePageDto").(*dto.UpdateStatusPageDto)

	// statusPage

	return context.JSON(123, "")
}
