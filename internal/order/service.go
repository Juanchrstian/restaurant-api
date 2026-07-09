package order

import "context"

type Service interface {
	CreateOrder(
		ctx context.Context,
	) (*Order, error)

	AddItem(
		ctx context.Context,
		orderID string,
		request AddOrderItemRequest,
	) (*OrderItem, error)

	GetOrder(
		ctx context.Context,
		id string,
	) (*Order, error)

	UpdateItem(
		ctx context.Context,
		orderID string,
		itemID string,
		request UpdateOrderItemRequest,
	) (*OrderItem, error)

	RemoveItem(
		ctx context.Context,
		orderID string,
		itemID string,
	) error
}
