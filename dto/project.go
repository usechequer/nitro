package dto

import "github.com/google/uuid"

type CreateProjectDto struct {
	UserUuid    uuid.UUID `json:"user_uuid" validate:"required"`
	Name        string    `json:"name" validate:"required,min=3"`
	Url         string    `json:"url" validate:"required"`
	Description string    `json:"description"`
}

type UpdateProjectDto struct {
	Uuid                      uuid.UUID `json:"uuid" validate:"required"`
	Name                      string    `json:"name" validate:"min=3"`
	Description               string    `json:"description"`
	Url                       string    `json:"url"`
	Is_notifications_silenced uint      `json:"is_notifications_silenced"`
}
