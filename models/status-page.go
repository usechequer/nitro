package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type StatusPage struct {
	ID          uint            `gorm:"primaryKey;not null" json:"-"`
	Uuid        uuid.UUID       `gorm:"not null" json:"uuid"`
	Title       string          `gorm:"type:varchar(191)" json:"title"`
	Description string          `gorm:"type:text" json:"description"`
	config      *datatypes.JSON `json:"config"`
	CreatedAt   time.Time       `gorm:"not null" json:"-"`
	UpdatedAt   time.Time       `gorm:"not null" json:"-"`
	ProjectID   uint            `gorm:"not null" json:"-"`
}

func (statusPage *StatusPage) BeforeCreate(transaction *gorm.DB) (err error) {
	statusPage.Uuid = uuid.New()
	return
}
