package menu

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	WithTransaction(
		tx *gorm.DB,
	) Repository

	GetAll(ctx context.Context, filter MenuFilter) ([]Menu, error)

	GetByID(ctx context.Context, id string) (*Menu, error)

	Create(ctx context.Context, menu *Menu) error

	Update(ctx context.Context, menu *Menu) error

	Delete(ctx context.Context, menu *Menu) error
}
