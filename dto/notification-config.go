package dto

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type CreateNotificationConfigDto struct {
	ProjectUuid uuid.UUID      `param:"project_uuid"`
	Config      datatypes.JSON `json:"config"`
}

type UpdateNotificationConfigDto struct {
	ProjectUuid uuid.UUID      `param:"project_uuid"`
	Uuid        uuid.UUID      `param:"uuid"`
	Config      datatypes.JSON `json:"config"`
}
