package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID                        uint           `gorm:"primaryKey;not null" json:"-"`
	Uuid                      uuid.UUID      `gorm:"index:projects_by_uuid;type:varchar(36);not null" json:"uuid"`
	Name                      string         `gorm:"not null;type:varchar(191)" json:"name"`
	Description               *string        `gorm:"type:text" json:"description"`
	Domain                    string         `gorm:"type:varchar(100)" json:"domain"`
	Logo                      *string        `gorm:"type:varchar(191)" json:"logo"`
	Is_notifications_silenced uint           `gorm:"not null" json:"is_notifications_silenced"`
	CreatedAt                 time.Time      `gorm:"not null" json:"-"`
	UpdatedAt                 time.Time      `gorm:"not null" json:"-"`
	DeletedAt                 gorm.DeletedAt `json:"-"`
	NotificationConfig        NotificationConfig
	StatusPage                StatusPage
}

func (project *Project) BeforeCreate(transaction *gorm.DB) (err error) {
	project.Uuid = uuid.New()
	return
}
