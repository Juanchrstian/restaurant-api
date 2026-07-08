package session

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Status Status `gorm:"size:20;not null"`

	OpenedBy string `gorm:"size:100;not null"`

	OpeningCash int64 `gorm:"type:numeric(12,2);not null"`

	ClosingCash *int64 `gorm:"type:numeric(12,2)"`

	OpenedAt time.Time `gorm:"not null"`

	ClosedAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
