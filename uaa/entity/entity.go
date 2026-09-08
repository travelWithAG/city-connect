package entity

import (
	"connect/core/shared-model"

	"gorm.io/datatypes"
)

type UserType string

const (
	AdminUserType UserType = "admin"
	NormalUserType UserType = "normal"
)

// User represents internal users, owners, and system administrators.
type User struct {
	sharedmodel.BaseEntity
	Email        string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash string         `gorm:"type:varchar(255);not null" json:"-"`
	FirstName    string         `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName     string         `gorm:"type:varchar(100);not null" json:"last_name"`
	PhoneNumber  string         `gorm:"type:varchar(20);index" json:"phone_number"`
	UserType 	 UserType	`gorm:"type:varchar(50);not null" json:"user_type"`
	Role         string         `gorm:"type:varchar(50);not null;default:'OWNER'" json:"role"` // OWNER, ADMIN, STAFF
	Metadata     datatypes.JSON `gorm:"type:jsonb;default:'{}'" json:"metadata,omitempty"`     // e.g., preferences, bank payout details
	// Properties   []Property     `gorm:"foreignKey:OwnerID" json:"properties,omitempty"`
}