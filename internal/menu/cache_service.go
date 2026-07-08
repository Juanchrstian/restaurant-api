package menu

import (
	"context"
	"encoding/json"
)

func (s *service) getMenusFromCache(
	ctx context.Context,
	cacheKey string,
) ([]Menu, bool, error) {

	data, err := s.cache.Get(
		ctx,
		cacheKey,
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
	cacheKey string,
	menus []Menu,
) {

	data, err := json.Marshal(menus)

	if err != nil {
		return
	}

	_ = s.cache.Set(
		ctx,
		cacheKey,
		data,
		MenuCacheTTL,
	)
}
