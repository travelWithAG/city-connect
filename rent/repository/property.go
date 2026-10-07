package repository

import (
	"connect/rent/entity"
	"connect/rent/handler"

	"gorm.io/gorm"
)

type PropertyRepository interface {
	CreateProperty(property *entity.Property) error
	// GetPropertyById(id uint) entity.Property
	// GetPropertiesByOwnerID(ownerID uint) ([]entity.Property, error)
}

type newProperty struct {
	db *gorm.DB
}

func NewPropertyRepo(db *gorm.DB) PropertyRepository  {
	return &newProperty{db}
}

func (pr *newProperty) CreateProperty(p *entity.Property) error {
	if err := pr.db.Create(p).Error; err != nil {
		return handler.ExceptionError()
	}
	return nil
}

func (pr *newProperty) CreateRoom(p entity.Room) error {
	if err := pr.db.Create("Property").Error; err != nil {
		return handler.ExceptionError()
	}
	return nil
}