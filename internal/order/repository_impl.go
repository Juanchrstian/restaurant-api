package order

import (
	"context"
	"fmt"

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

func (r *repository) CreateItem(
	ctx context.Context,
	item *OrderItem,
) error {

	return r.db.
		WithContext(ctx).
		Create(item).
		Error
}

func (r *repository) WithTransaction(
	tx *gorm.DB,
) Repository {

	return &repository{
		db: tx,
	}

}

func (r *repository) DB() *gorm.DB {
	return r.db
}

func (r *repository) GetDetail(
	ctx context.Context,
	id string,
) (*Order, error) {

	var order Order

	err := r.db.
		WithContext(ctx).
		Preload("Items").
		Preload("Items.Menu").
		Where("id = ?", id).
		First(&order).
		Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *repository) GetItemByID(
	ctx context.Context,
	itemID string,
) (*OrderItem, error) {

	var item OrderItem

	err := r.db.
		WithContext(ctx).
		First(&item, "id = ?", itemID).
		Error

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, sharederrors.ErrOrderItemNotFound
		}

		return nil, err
	}

	return &item, nil
}

func (r *repository) GetItemByMenu(
	ctx context.Context,
	orderID string,
	menuID string,
) (*OrderItem, error) {

	var item OrderItem

	err := r.db.
		WithContext(ctx).
		Where(
			"order_id = ? AND menu_id = ?",
			orderID,
			menuID,
		).
		First(&item).
		Error

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &item, nil
}

func (r *repository) UpdateItem(
	ctx context.Context,
	item *OrderItem,
) error {

	return r.db.
		WithContext(ctx).
		Save(item).
		Error
}

func (r *repository) DeleteItem(
	ctx context.Context,
	item *OrderItem,
) error {

	return r.db.
		WithContext(ctx).
		Delete(item).
		Error
}

func (r *repository) GetItems(
	ctx context.Context,
	orderID string,
) ([]OrderItem, error) {

	var items []OrderItem

	err := r.db.
		WithContext(ctx).
		Where("order_id = ?", orderID).
		Find(&items).
		Error

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *repository) GetAll(
	ctx context.Context,
	request GetOrdersRequest,
) ([]Order, error) {

	var orders []Order

	db := r.db.WithContext(ctx)

	if request.Status != "" {
		db = db.Where(
			"status = ?",
			request.Status,
		)
	}

	if request.PaymentMethod != "" {
		db = db.Where(
			"payment_method = ?",
			request.PaymentMethod,
		)
	}

	if request.SessionID != "" {
		db = db.Where(
			"session_id = ?",
			request.SessionID,
		)
	}

	orderBy := fmt.Sprintf(
		"%s %s",
		request.Sort,
		request.Order,
	)

	offset := (request.Page - 1) * request.Limit

	err := db.
		Preload("Items").
		Order(orderBy).
		Offset(offset).
		Limit(request.Limit).
		Find(&orders).Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}
