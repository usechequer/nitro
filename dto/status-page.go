package dto

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type CreateStatusPageDto struct {
	ProjectUuid uuid.UUID       `param:"project_uuid"`
	Title       *string         `json:"title"`
	Description *string         `json:"description"`
	Metadata    *datatypes.JSON `json:"metadata"`
	Config      datatypes.JSON  `json:"config" validate:"required"`
}

type UpdateStatusPageDto struct {
	CreateStatusPageDto
	Uuid   uuid.UUID       `param:"uuid"`
	Config *datatypes.JSON `json:"config"`
}
