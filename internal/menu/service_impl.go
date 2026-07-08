package menu

import (
	"context"

	"github.com/juanchrstian/restaurant-api/internal/shared/cache"
	sharedcache "github.com/juanchrstian/restaurant-api/internal/shared/cache"
)

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
	filter MenuFilter,
) ([]Menu, error) {

	return s.repository.GetAll(
		ctx,
		filter,
	)

	// menus, found, err := s.getMenusFromCache(ctx)

	// if err == nil && found {

	// 	log.Println("CACHE HIT")

	// 	return menus, nil

	// }

	// log.Println("CACHE MISS")

	// menus, err = s.repository.GetAll(ctx, filter)

	// if err != nil {
	// 	return nil, err
	// }

	// s.saveMenusToCache(
	// 	ctx,
	// 	menus,
	// )

	// return menus, nil

}

func (s *service) GetMenu(
	ctx context.Context,
	id string,
) (*Menu, error) {

	return s.repository.GetByID(ctx, id)

}

func (s *service) CreateMenu(
	ctx context.Context,
	request CreateMenuRequest,
) (*Menu, error) {

	menu := request.ToModel()

	if err := s.repository.Create(ctx, &menu); err != nil {
		return nil, err
	}

	// Cache Invalidation
	_ = s.cache.Delete(
		ctx,
		sharedcache.MenuListKey,
	)

	return &menu, nil
}

func (s *service) UpdateMenu(
	ctx context.Context,
	id string,
	request UpdateMenuRequest,
) (*Menu, error) {

	menu, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	request.Apply(menu)

	if err := s.repository.Update(ctx, menu); err != nil {
		return nil, err
	}

	_ = s.cache.Delete(
		ctx,
		sharedcache.MenuListKey,
	)

	return menu, nil
}

func (s *service) DeleteMenu(
	ctx context.Context,
	id string,
) error {

	menu, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := s.repository.Delete(ctx, menu); err != nil {
		return err
	}

	_ = s.cache.Delete(
		ctx,
		sharedcache.MenuListKey,
	)

	return nil
}
