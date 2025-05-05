package dto

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type CreateNotificationConfig struct {
	ProjectUuid uuid.UUID      `param:"project_uuid"`
	Config      datatypes.JSON `json:"config"`
}
