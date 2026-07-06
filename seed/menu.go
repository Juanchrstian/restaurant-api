package seed

import (
	"github.com/google/uuid"
	"github.com/juanchrstian/restaurant-api/internal/menu"
	"gorm.io/gorm"
)

func SeedMenus(db *gorm.DB) error {

	var count int64

	db.Model(&menu.Menu{}).Count(&count)

	if count > 0 {
		return nil
	}

	menus := []menu.Menu{
		{
			ID:          uuid.New(),
			Name:        "Nasi Goreng",
			Description: "Indonesian fried rice",
			Price:       25000,
			Stock:       20,
			Available:   true,
		},
		{
			ID:          uuid.New(),
			Name:        "Mie Goreng",
			Description: "Fried noodles",
			Price:       22000,
			Stock:       20,
			Available:   true,
		},
		{
			ID:          uuid.New(),
			Name:        "Ayam Geprek",
			Description: "Spicy fried chicken",
			Price:       30000,
			Stock:       15,
			Available:   true,
		},
		{
			ID:          uuid.New(),
			Name:        "Es Teh",
			Description: "Iced Tea",
			Price:       7000,
			Stock:       50,
			Available:   true,
		},
		{
			ID:          uuid.New(),
			Name:        "Es Jeruk",
			Description: "Orange Juice",
			Price:       10000,
			Stock:       30,
			Available:   true,
		},
	}

	return db.Create(&menus).Error
}