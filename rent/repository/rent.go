package repository

import (
	"connect/rent/entity"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RentRepository interface {
	RentProperty(entity.RentRecord, *entity.Customer) error
}

type rentRepo struct {
	db *gorm.DB
}

func NewRentRepository(db *gorm.DB) RentRepository  {
	return &rentRepo{ db: db }
}

func (rr *rentRepo) RentProperty(rent entity.RentRecord, customer *entity.Customer) error {
	return rr.db.Transaction(func(tx *gorm.DB) error {
		// If CustomerID is not provided, create the customer first
		if rent.CustomerID == uuid.Nil {
			if customer == nil {
				return errors.New("customer data is required when customer ID is not provided")
			}

			if err := tx.Create(customer).Error; err != nil {
				return err
			}

			// Assign the newly generated customer ID to the rent record
			rent.CustomerID = customer.ID
		}

		// Create the rent record
		if err := tx.Create(&rent).Error; err != nil {
			return err
		}

		return nil
	})
}