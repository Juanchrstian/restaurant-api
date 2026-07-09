package errors

import "errors"

var (
	ErrMenuNotFound      = errors.New("Menu not found")
	ErrMenuUnavailable   = errors.New("menu unavailable")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrOrderItemNotFound = errors.New("order item not found")
)
