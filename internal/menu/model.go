package menu

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Menu struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string    `gorm:"size:100;not null"`
	Description string    `gorm:"type:text"`
	Price       int64     `gorm:"type:numeric(12,2);not null"`
	Stock       int       `gorm:"not null"`
	Available   bool      `gorm:"default:true"`
	ImageURL    string    `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
