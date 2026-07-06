package menu

import (
	"context"
	"encoding/json"

	sharedcache "github.com/juanchrstian/restaurant-api/internal/shared/cache"
)

func (s *service) getMenusFromCache(
	ctx context.Context,
) ([]Menu, bool, error) {

	data, err := s.cache.Get(
		ctx,
		sharedcache.MenuListKey,
	)

	if err != nil {
		return nil, false, nil
	}

	var menus []Menu

	if err := json.Unmarshal(
		[]byte(data),
		&menus,
	); err != nil {

		return nil, false, err
	}

	return menus, true, nil
}

func (s *service) saveMenusToCache(
	ctx context.Context,
	menus []Menu,
) {

	data, err := json.Marshal(menus)

	if err != nil {
		return
	}

	_ = s.cache.Set(
		ctx,
		sharedcache.MenuListKey,
		data,
		MenuCacheTTL,
	)
}