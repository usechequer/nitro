package dto

import "github.com/google/uuid"

type CreateProjectDto struct {
	UserUuid    uuid.UUID `json:"user_uuid" validate:"required"`
	Name        string    `json:"name" validate:"required,min=3"`
	Url         string    `json:"url" validate:"required"`
	Description string    `json:"description"`
}
