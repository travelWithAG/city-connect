package sharedmodel

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseEntity provides common attributes for enterprise auditability.
type BaseEntity struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime;index" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Active    bool           `gorm:"default:true" json:"active"`
}