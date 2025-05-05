package controllers

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"nitro/dto"
	"nitro/models"
	"nitro/utilities"
	"os"

	cloudinaryV2 "github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

var logoErrorMessage string = "There was a problem uploading the logo"

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
	project := context.Get("project").(models.Project)
	updateProjectDto := context.Get("updateProjectDto").(*dto.UpdateProjectDto)

	updateProjectField := func(field string, param string) string {
		if len(param) > 0 {
			return param
		}

		return field
	}

	updateProjectDescription := func(description string, descriptionParam string) *string {
		var descriptionPointer *string

		if len(descriptionParam) > 0 {
			descriptionPointer = &descriptionParam
		} else {
			descriptionPointer = &description
		}

		return descriptionPointer
	}

	project.Name = updateProjectField(project.Name, updateProjectDto.Name)
	project.Url = updateProjectField(project.Url, updateProjectDto.Url)
	project.Description = updateProjectDescription(*project.Description, updateProjectDto.Description)

	if updateProjectDto.Is_notifications_silenced != nil {
		project.Is_notifications_silenced = *updateProjectDto.Is_notifications_silenced
	}

	database := utilities.GetDatabaseObject()

	logo, err := context.FormFile("logo")

	if err != nil {
		database.Save(&project)
		return context.JSON(http.StatusOK, project)
	}

	logoSrc, err := uploadLogo(logo, project.Uuid.String())

	if err != nil {
		return utilities.ThrowException(context, &utilities.Exception{StatusCode: http.StatusInternalServerError, Error: "INTERNAL_SERVER_ERROR", Message: logoErrorMessage})
	}

	project.Logo = &logoSrc
	database.Save(&project)

	return context.JSON(http.StatusOK, project)
}

func uploadLogo(logo *multipart.FileHeader, projectUuid string) (logoSrc string, err error) {
	cloudinary, err := cloudinaryV2.New()

	if err != nil {
		return "", errors.New(logoErrorMessage)
	}

	src, err := logo.Open()

	if err != nil {
		return "", errors.New(logoErrorMessage)
	}

	defer src.Close()

	uploadResult, err := cloudinary.Upload.Upload(context.Background(), src, uploader.UploadParams{Folder: fmt.Sprintf("%s/projects", os.Getenv("CLOUDINARY_FOLDER")), ResourceType: "image", PublicID: projectUuid, Overwrite: api.Bool(true)})

	if err != nil {
		return "", errors.New(logoErrorMessage)
	}

	return uploadResult.SecureURL, nil
}
