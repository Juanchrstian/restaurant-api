package menu

import "context"

var _ Service = (*service)(nil)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetMenus(
	ctx context.Context,
) ([]Menu, error) {

	return s.repository.GetAll(ctx)

}

func (s *service) GetMenu(
	ctx context.Context,
	id string,
) (*Menu, error) {

	return s.repository.GetByID(ctx, id)

}