package migrations

import (
	"fmt"

	"gorm.io/gorm"
	uaaEntity "connect/uaa/entity"
	rentEntity "connect/rent/entity"
)

func Run(db *gorm.DB) error {
    err := db.AutoMigrate(
        &uaaEntity.User{},
		&rentEntity.Property{},
		&rentEntity.Room{},
		&rentEntity.Customer{},
		&rentEntity.RentRecord{},
		&rentEntity.MaintenanceTicket{},
    )
    if err != nil {
        return fmt.Errorf("failed to run database auto-migrations: %w", err)
    }

    return nil
}