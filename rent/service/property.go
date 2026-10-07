package service

import (
	"connect/rent/entity"
	"connect/rent/repository"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
)

type PropertyInput struct {
	OwnerID     uuid.UUID    `json:"owner_id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Type        entity.PropertyType `json:"type"`

	// Address Details
	AddressLine1 string   `gorm:"type:varchar(255);not null" json:"address_line_1"`
	AddressLine2 string   `gorm:"type:varchar(255)" json:"address_line_2"`
	City         string   `gorm:"type:varchar(100);not null;index" json:"city"`
	State        string   `gorm:"type:varchar(100)" json:"state"`
	PostalCode   string   `gorm:"type:varchar(20)" json:"postal_code"`
	Country      string   `json:"country"`
	Latitude     *float64 `json:"latitude,omitempty"`
	Longitude    *float64 `json:"longitude,omitempty"`

	// Extensible Specs & Dynamic Amenities
	// Example: {"parking_spaces": 4, "amenities": ["generator", "water_tank", "security_guard"]}
	Features datatypes.JSON `json:"features,omitempty"`
}

type RoomInput struct {
	PropertyID uuid.UUID `json:"property_id"`
	RoomNumber string    `json:"room_number"`
	FloorLevel int       `json:"floor_level"`

	// Precise Financial Valuation
	BasePrice decimal.Decimal `json:"base_price"`
	IsOccupied bool           `json:"is_occupied"`

	// Room-specific amenities (e.g., {"air_conditioned": true, "dimensions_sqm": 25.5})
	Features datatypes.JSON `json:"features,omitempty"`
}

type PropertyService interface {
	CreateProperty(property PropertyInput) (*entity.Property, error)
	// CreateRoom(property RoomInput) (entity.Room, error)
}

type newProperty struct {
	propertyService repository.PropertyRepository
}

func NewPropertyService(ps repository.PropertyRepository) PropertyService {
	return &newProperty{propertyService: ps}
}

func (p *newProperty) CreateProperty(prop PropertyInput) (*entity.Property, error)  {
	if strings.TrimSpace(prop.Title) == "" {
		return nil, errors.New("property name is required")
	}

	fmt.Println("Property type submitted: ", prop.Type)

	// Validate Property Type
	switch prop.Type {
		case entity.TypeHouse, entity.TypeApartment, entity.TypeGuestHouse, entity.TypeShopFrame:
			// Valid type
		default:
			return nil, fmt.Errorf("invalid property type: %s", prop.Type)
	}

	property := &entity.Property{
		OwnerID:     prop.OwnerID,
		Title:       prop.Title,
		Description: prop.Description,
		Type: prop.Type,
		AddressLine1:     prop.AddressLine1,
		City:        prop.City,
	}

	if err := p.propertyService.CreateProperty(property); err != nil {
		return nil, fmt.Errorf("failed to create property: %w", err)
	}

	return property, nil
}