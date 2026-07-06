package menu

import "context"

type Repository interface {

	GetAll(ctx context.Context) ([]Menu, error)

	GetByID(ctx context.Context, id string) (*Menu, error)

}