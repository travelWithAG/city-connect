package repository

import (
	"connect/rent/entity"
	"fmt"

	"gorm.io/gorm"
)

type MaintanenceRepository interface {
	RequestMantainence(dto entity.MaintenanceTicket) error
}

type maintanenceRepo struct {
	db *gorm.DB
}

func NewMaintanenceRepository(db *gorm.DB) MaintanenceRepository {
	return &maintanenceRepo{db: db}
}

func (dbc *maintanenceRepo) RequestMantainence(dto entity.MaintenanceTicket) error {
	err := dbc.db.Create(dto).Error 
	if err != nil {
		return fmt.Errorf("Error occured")
	}
	return nil
}