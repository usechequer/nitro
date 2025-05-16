package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type StatusPage struct {
	ID          uint            `gorm:"primaryKey;not null" json:"-"`
	Uuid        uuid.UUID       `gorm:"index:status_pages_by_uuid;type:varchar(36);not null" json:"uuid"`
	ProjectID   uint            `gorm:"not null" json:"-"`
	Project     Project         `gorm:"constraint:OnUpdate:CASCADE,onDelete:RESTRICT" json:"-"`
	Title       *string         `gorm:"type:varchar(191)" json:"title"`
	Description *string         `gorm:"type:text" json:"description"`
	Metadata    *datatypes.JSON `json:"metadata"`
	Config      datatypes.JSON  `json:"config"`
	CreatedAt   time.Time       `gorm:"not null" json:"-"`
	UpdatedAt   time.Time       `gorm:"not null" json:"-"`
}

func (statusPage *StatusPage) BeforeCreate(transaction *gorm.DB) (err error) {
	statusPage.Uuid = uuid.New()
	return
}
