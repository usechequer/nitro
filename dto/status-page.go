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
	ProjectUuid uuid.UUID       `param:"project_uuid"`
	Uuid        uuid.UUID       `param:"uuid"`
	Title       *string         `json:"title"`
	Description *string         `json:"description"`
	Metadata    *datatypes.JSON `json:"metadata"`
	Config      *datatypes.JSON `json:"config"`
}
