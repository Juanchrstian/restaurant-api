package menu

import "context"

type Service interface {
	GetMenus(ctx context.Context, filter MenuFilter) ([]Menu, error)

	GetMenu(ctx context.Context, id string) (*Menu, error)

	CreateMenu(ctx context.Context, request CreateMenuRequest) (*Menu, error)

	UpdateMenu(ctx context.Context, id string, request UpdateMenuRequest) (*Menu, error)

	DeleteMenu(ctx context.Context, id string) error
}
