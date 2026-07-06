package menu

import "context"

type Service interface {
	GetMenus(ctx context.Context) ([]Menu, error)

	GetMenu(ctx context.Context, id string) (*Menu, error)

	CreateMenu(ctx context.Context, request CreateMenuRequest) (*Menu, error)
}
