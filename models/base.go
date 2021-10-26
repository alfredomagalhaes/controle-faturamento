package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID  `gorm:"type:varchar(36);primary_key;";json:"id"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewV4()

	base.ID = uuid

	return nil
}
