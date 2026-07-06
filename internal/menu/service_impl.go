package menu

import "context"

import "log"

import "github.com/juanchrstian/restaurant-api/internal/shared/cache"

var _ Service = (*service)(nil)

type service struct {
	
	repository Repository

	cache cache.Cache

}

func NewService(
	
	repository Repository,
	
	cache cache.Cache,
	
	) Service {

	return &service{

		repository: repository,

		cache: cache,
	}
}

func (s *service) GetMenus(
		ctx context.Context,
	) ([]Menu, error) {

	menus, found, err := s.getMenusFromCache(ctx)

	if err == nil && found {

		log.Println("CACHE HIT")

		return menus, nil

	}

	log.Println("CACHE MISS")

	menus, err = s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	s.saveMenusToCache(
		ctx,
		menus,
	)

	return menus, nil

}

func (s *service) GetMenu(
	ctx context.Context,
	id string,
) (*Menu, error) {

	return s.repository.GetByID(ctx, id)

}