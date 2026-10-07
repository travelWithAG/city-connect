package service

import (
	"connect/rent/entity"
	"connect/rent/repository"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
)

type RentDto struct {
	CustomerID uuid.UUID `json:"customer_id"`
	Email        string         `json:"email"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	PhoneNumber  string         `json:"phone_number"`
	IDType       string         `json:"id_type"` // NIDA, Passport, Driver License
	IDNumber     string         `json:"id_number"`
	Metadata     datatypes.JSON `json:"metadata,omitempty"` // Emergency contact details, KYC verification flags

	PropertyID uuid.UUID  `json:"property_id"`
	RoomID     *uuid.UUID `json:"room_id,omitempty"`

	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`

	// Financial Amounts using fixed precision
	RentAmount    decimal.Decimal `json:"rent_amount"`
	DepositAmount decimal.Decimal `json:"deposit_amount"`

	PaymentCycle string         ` json:"payment_cycle"` // MONTHLY, QUARTERLY, ANNUALLY
	Status       string         `json:"status"`   // ACTIVE, TERMINATED, EXPIRED
	Terms  datatypes.JSON `json:"terms,omitempty"`
}

type RentRecordService interface {
	RentProperty(dto RentDto) (*entity.RentRecord, error)
}

type rentService struct {
	rentRepo repository.RentRepository
}

func (rs *rentService) NewRentRecordService(repo repository.RentRepository) RentRecordService {
	return &rentService{rentRepo: repo}
}

func (rs *rentService) RentProperty(dto RentDto) (*entity.RentRecord, error) {

	if dto.CustomerID == uuid.Nil{
		var errs []string

		if strings.TrimSpace(dto.Email) == ""{
			errs = append(errs, "Email is Required.")
		}

		if strings.TrimSpace(dto.LastName) == ""{
			errs = append(errs, "Last Name is Required.")
		}

		if strings.TrimSpace(dto.PhoneNumber) == ""{
			errs = append(errs, "Phone Number is Required.")
		}

		if strings.TrimSpace(dto.FirstName) == ""{
			errs = append(errs, "First Name is Required.")
		}

		if (dto.PropertyID == uuid.Nil && dto.RoomID == nil) {
			errs = append(errs, "Property For rent is Required.")
		}

		return nil, errors.New(strings.Join(errs, ", "))
	}
	rentRec := &entity.RentRecord{
		CustomerID: dto.CustomerID,				
		PropertyID: dto.PropertyID,
		RoomID: dto.RoomID,
		StartDate: dto.StartDate,
		EndDate: dto.EndDate,
		RentAmount: dto.RentAmount,
		DepositAmount: dto.DepositAmount,
		PaymentCycle: dto.PaymentCycle,
		Status: dto.Status,
		Terms: dto.Terms,
	}

	customer := &entity.Customer{
		Email: dto.Email,
		FirstName: dto.FirstName,
		LastName: dto.LastName,
		PhoneNumber: dto.PhoneNumber,
		IDType: dto.IDNumber,
		IDNumber: dto.IDNumber, 
		Metadata: dto.Metadata,
	}

	err := rs.rentRepo.RentProperty(*rentRec, customer)

	if(err != nil){
		return nil, fmt.Errorf("Error Occured %v", err)
	}

	return rentRec, nil
}