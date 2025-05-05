package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type NotificationConfig struct {
	ID        uint           `gorm:"primaryKey;not null" json:"-"`
	Uuid      uuid.UUID      `gorm:"index:notification_configs_by_uuid;type:varchar(36);not null" json:"uuid"`
	ProjectID uint           `gorm:"not null" json:"-"`
	Config    datatypes.JSON `gorm:"not null" json:"config"`
	CreatedAt time.Time      `gorm:"not null" json:"-"`
	UpdatedAt time.Time      `gorm:"not null" json:"-"`
}

func (notificationConfig *NotificationConfig) BeforeCreate(transaction *gorm.DB) (err error) {
	notificationConfig.Uuid = uuid.New()
	return
}
