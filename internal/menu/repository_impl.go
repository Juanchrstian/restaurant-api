package menu

import (
	"context"
	"strings"

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
	filter MenuFilter,
) ([]Menu, error) {

	var menus []Menu

	query := r.db.
		WithContext(ctx).
		Model(&Menu{})

	if filter.Available != nil {
		query = query.Where(
			"available = ?",
			*filter.Available,
		)
	}

	if filter.Search != "" {
		query = query.Where(
			"LOWER(name) LIKE ?",
			"%"+strings.ToLower(filter.Search)+"%",
		)
	}

	query = query.Order(filter.OrderClause()).
		Order("name ASC")

	offset := (filter.Page - 1) * filter.Limit

	query = query.
		Offset(offset).
		Limit(filter.Limit)

	if err := query.
		Debug().
		Find(&menus).
		Error; err != nil {

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

func (r *repository) Create(
	ctx context.Context,
	menu *Menu,
) error {

	return r.db.
		WithContext(ctx).
		Create(menu).
		Error
}

func (r *repository) Update(
	ctx context.Context,
	menu *Menu,
) error {

	return r.db.
		WithContext(ctx).
		Save(menu).
		Error
}

func (r *repository) Delete(
	ctx context.Context,
	menu *Menu,
) error {

	return r.db.
		WithContext(ctx).
		Delete(menu).
		Error
}
