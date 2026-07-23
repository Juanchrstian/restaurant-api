package order

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	DB() *gorm.DB

	WithTransaction(
		tx *gorm.DB,
	) Repository

	Create(
		ctx context.Context,
		order *Order,
	) error

	GetByID(
		ctx context.Context,
		id string,
	) (*Order, error)

	Update(
		ctx context.Context,
		order *Order,
	) error

	CreateItem(
		ctx context.Context,
		item *OrderItem,
	) error

	GetDetail(
		ctx context.Context,
		id string,
	) (*Order, error)

	GetItemByID(
		ctx context.Context,
		itemID string,
	) (*OrderItem, error)

	GetItemByMenu(
		ctx context.Context,
		orderID string,
		menuID string,
	) (*OrderItem, error)

	UpdateItem(
		ctx context.Context,
		item *OrderItem,
	) error

	DeleteItem(
		ctx context.Context,
		item *OrderItem,
	) error

	GetItems(
		ctx context.Context,
		orderID string,
	) ([]OrderItem, error)

	GetAll(
		ctx context.Context,
		request GetOrdersRequest,
	) ([]Order, error)
}
