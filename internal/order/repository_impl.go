package order

import (
	"context"

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

func (r *repository) Create(
	ctx context.Context,
	order *Order,
) error {

	return r.db.
		WithContext(ctx).
		Create(order).
		Error
}

func (r *repository) GetByID(
	ctx context.Context,
	id string,
) (*Order, error) {

	var order Order

	err := r.db.
		WithContext(ctx).
		Where("id = ?", id).
		First(&order).
		Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *repository) Update(
	ctx context.Context,
	order *Order,
) error {

	return r.db.
		WithContext(ctx).
		Save(order).
		Error
}
