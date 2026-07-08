package order

import "context"

type Service interface {
	CreateOrder(
		ctx context.Context,
	) (*Order, error)
}
