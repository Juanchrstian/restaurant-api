package menu

import (
	"context"

	sharederrors "github.com/juanchrstian/restaurant-api/internal/shared/errors"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(
	ctx context.Context,
) ([]Menu, error) {

	var menus []Menu

	err := r.db.
		WithContext(ctx).
		Where("available = ?", true).
		Order("name ASC").
		Find(&menus).
		Error

	if err != nil {
		return nil, err
	}

	return menus, nil
}

func (r *repository) GetByID(
	ctx context.Context,
	id string,
) (*Menu, error) {

	var menu Menu

	err := r.db.
		WithContext(ctx).
		Where("id = ?", id).
		First(&menu).
		Error

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, sharederrors.ErrMenuNotFound
		}

		return nil, err
	}

	return &menu, nil
}