package order

import (
	"time"

	"github.com/google/uuid"
	"github.com/juanchrstian/restaurant-api/internal/menu"
	"gorm.io/gorm"
)

type Order struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	SessionID uuid.UUID `gorm:"type:uuid;not null"`

	OrderNumber string `gorm:"size:50;uniqueIndex;not null"`

	Status OrderStatus `gorm:"size:20;not null"`

	TotalAmount int64 `gorm:"not null;default:0"`

	CreatedAt time.Time
	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt `gorm:"index"`

	Items []OrderItem `gorm:"foreignKey:OrderID"`

	// =========================
	// PAYMENT
	// =========================

	PaymentMethod *PaymentMethod

	PaidAmount *int64

	ChangeAmount *int64

	PaidAt *time.Time
}

type OrderItem struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	OrderID uuid.UUID `gorm:"type:uuid;not null"`

	MenuID uuid.UUID `gorm:"type:uuid;not null"`

	Quantity int `gorm:"not null"`

	Price int64 `gorm:"not null"`

	Subtotal int64 `gorm:"not null"`

	CreatedAt time.Time

	Menu menu.Menu `gorm:"foreignKey:MenuID"`
}
