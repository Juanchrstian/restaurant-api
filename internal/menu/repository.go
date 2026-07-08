package menu

import "context"

type Repository interface {
	GetAll(ctx context.Context, filter MenuFilter) ([]Menu, error)

	GetByID(ctx context.Context, id string) (*Menu, error)

	Create(ctx context.Context, menu *Menu) error

	Update(ctx context.Context, menu *Menu) error

	Delete(ctx context.Context, menu *Menu) error
}
