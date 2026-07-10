package errors

import "errors"

var (
	ErrOrderItemNotFound = errors.New("order item not found")

	ErrOrderNotFound = errors.New("order not found")

	ErrOrderAlreadyPaid = errors.New("order already paid")

	ErrInsufficientPayment = errors.New("insufficient payment")

	ErrEmptyOrder = errors.New("empty order")
)
