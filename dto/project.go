package dto

import "github.com/google/uuid"

type CreateProjectDto struct {
	UserUuid    uuid.UUID `json:"user_uuid" validate:"required"`
	Name        string    `json:"name" validate:"required,min=3"`
	Url         string    `json:"url" validate:"required"`
	Description string    `json:"description"`
}

type UpdateProjectDto struct {
	Uuid                      uuid.UUID `param:"uuid"`
	Name                      string    `form:"name"`
	Description               string    `form:"description"`
	Url                       string    `form:"url"`
	Is_notifications_silenced *uint     `form:"is_notifications_silenced"`
}
