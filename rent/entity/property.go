package entity

import (
	sharedmodel "connect/core/shared-model"
	uaaEntity "connect/uaa/entity"
	"time"

	"gorm.io/datatypes"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)
type PropertyType string

const (
	TypeHouse      PropertyType = "HOUSE"
	TypeApartment  PropertyType = "APARTMENT"
	TypeGuestHouse PropertyType = "GUEST_HOUSE"
	TypeShopFrame  PropertyType = "SHOP_FRAME"
)

// Property defines high-level real estate assets owned by a User.
type Property struct {
	sharedmodel.BaseEntity
	OwnerID     uuid.UUID    `gorm:"type:uuid;not null;index" json:"owner_id"`
	Owner       uaaEntity.User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"owner,omitempty"`
	Title       string       `gorm:"type:varchar(200);not null" json:"title"`
	Description string       `gorm:"type:text" json:"description"`
	Type        PropertyType `gorm:"type:varchar(50);not null;index" json:"type"`

	// Address Details
	AddressLine1 string   `gorm:"type:varchar(255);not null" json:"address_line_1"`
	AddressLine2 string   `gorm:"type:varchar(255)" json:"address_line_2"`
	City         string   `gorm:"type:varchar(100);not null;index" json:"city"`
	State        string   `gorm:"type:varchar(100)" json:"state"`
	PostalCode   string   `gorm:"type:varchar(20)" json:"postal_code"`
	Country      string   `gorm:"type:varchar(100);not null;default:'Tanzania'" json:"country"`
	Latitude     *float64 `gorm:"type:decimal(10,8)" json:"latitude,omitempty"`
	Longitude    *float64 `gorm:"type:decimal(11,8)" json:"longitude,omitempty"`

	// Extensible Specs & Dynamic Amenities
	// Example: {"parking_spaces": 4, "amenities": ["generator", "water_tank", "security_guard"]}
	Features datatypes.JSON `gorm:"type:jsonb;default:'{}'" json:"features,omitempty"`

	// Relationships
	Rooms       []Room       `gorm:"foreignKey:PropertyID" json:"rooms,omitempty"`
	RentRecords []RentRecord `gorm:"foreignKey:PropertyID" json:"rent_records,omitempty"`
}

// Room handles sub-unit rentals (e.g., individual room in a house or shop frame unit).
type Room struct {
	sharedmodel.BaseEntity
	PropertyID uuid.UUID `gorm:"type:uuid;not null;index" json:"property_id"`
	Property   Property  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"property,omitempty"`
	RoomNumber string    `gorm:"type:varchar(50);not null" json:"room_number"`
	FloorLevel int       `gorm:"default:0" json:"floor_level"`

	// Precise Financial Valuation
	BasePrice decimal.Decimal `gorm:"type:decimal(12,2);not null" json:"base_price"`
	IsOccupied bool           `gorm:"default:false;index" json:"is_occupied"`

	// Room-specific amenities (e.g., {"air_conditioned": true, "dimensions_sqm": 25.5})
	Features datatypes.JSON `gorm:"type:jsonb;default:'{}'" json:"features,omitempty"`

	RentRecords []RentRecord `gorm:"foreignKey:RoomID" json:"rent_records,omitempty"`
}

// Customer represents tenant accounts interacting with the portal.
type Customer struct {
	sharedmodel.BaseEntity
	Email        string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash string         `gorm:"type:varchar(255);not null" json:"-"`
	FirstName    string         `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName     string         `gorm:"type:varchar(100);not null" json:"last_name"`
	PhoneNumber  string         `gorm:"type:varchar(20);not null;index" json:"phone_number"`
	IDType       string         `gorm:"type:varchar(50)" json:"id_type"` // NIDA, Passport, Driver License
	IDNumber     string         `gorm:"type:varchar(100);index" json:"id_number"`
	Metadata     datatypes.JSON `gorm:"type:jsonb;default:'{}'" json:"metadata,omitempty"` // Emergency contact details, KYC verification flags

	RentRecords []RentRecord        `gorm:"foreignKey:CustomerID" json:"rent_records,omitempty"`
	Tickets     []MaintenanceTicket `gorm:"foreignKey:CustomerID" json:"tickets,omitempty"`
}

// RentRecord tracks contract timelines, financial amounts, and deposits.
type RentRecord struct {
	sharedmodel.BaseEntity
	CustomerID uuid.UUID `gorm:"type:uuid;not null;index" json:"customer_id"`
	Customer   Customer  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"customer,omitempty"`

	PropertyID uuid.UUID  `gorm:"type:uuid;not null;index" json:"property_id"`
	Property   Property   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"property,omitempty"`
	RoomID     *uuid.UUID `gorm:"type:uuid;index" json:"room_id,omitempty"`
	Room       *Room      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"room,omitempty"`

	StartDate time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate   time.Time `gorm:"type:date;not null;index" json:"end_date"`

	// Financial Amounts using fixed precision
	RentAmount    decimal.Decimal `gorm:"type:decimal(12,2);not null" json:"rent_amount"`
	DepositAmount decimal.Decimal `gorm:"type:decimal(12,2);default:0.00" json:"deposit_amount"`

	PaymentCycle string         `gorm:"type:varchar(20);default:'MONTHLY'" json:"payment_cycle"` // MONTHLY, QUARTERLY, ANNUALLY
	Status       string         `gorm:"type:varchar(30);default:'ACTIVE';index" json:"status"`   // ACTIVE, TERMINATED, EXPIRED
	Terms        datatypes.JSON `gorm:"type:jsonb;default:'{}'" json:"terms,omitempty"`          // Dynamic lease terms, penalty clauses
}

type MaintenanceCategory string

const (
	CategoryMaintenance MaintenanceCategory = "MAINTENANCE"
	CategoryEmergency   MaintenanceCategory = "EMERGENCY"
)

// MaintenanceTicket tracks routine maintenance and urgent emergency reports.
type MaintenanceTicket struct {
	sharedmodel.BaseEntity
	CustomerID uuid.UUID `gorm:"type:uuid;not null;index" json:"customer_id"`
	Customer   Customer  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"customer,omitempty"`
	PropertyID uuid.UUID `gorm:"type:uuid;not null;index" json:"property_id"`
	Property   Property  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"property,omitempty"`
	RoomID     *uuid.UUID `gorm:"type:uuid;index" json:"room_id,omitempty"`
	Room       *Room     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"room,omitempty"`

	Category MaintenanceCategory `gorm:"type:varchar(30);not null;default:'MAINTENANCE';index" json:"category"`
	Priority string              `gorm:"type:varchar(20);default:'MEDIUM';index" json:"priority"` // LOW, MEDIUM, HIGH, CRITICAL
	Status   string              `gorm:"type:varchar(30);default:'OPEN';index" json:"status"`     // OPEN, IN_PROGRESS, RESOLVED, CLOSED
	Subject  string              `gorm:"type:varchar(200);not null" json:"subject"`
	Description string           `gorm:"type:text;not null" json:"description"`

	// Estimated and Actual Repair Costs
	EstimatedCost decimal.Decimal  `gorm:"type:decimal(12,2);default:0.00" json:"estimated_cost"`
	ActualCost    decimal.Decimal  `gorm:"type:decimal(12,2);default:0.00" json:"actual_cost"`

	// Structured metadata for issue media, contractor notes, hardware tags
	// Example: {"photo_urls": ["s3://..."], "vendor_assigned": "Alpha Electric"}
	Details    datatypes.JSON `gorm:"type:jsonb;default:'{}'" json:"details,omitempty"`
	ResolvedAt *time.Time     `json:"resolved_at,omitempty"`
}