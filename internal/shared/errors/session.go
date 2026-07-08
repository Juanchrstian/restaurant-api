package errors

import "errors"

var (
	ErrSessionNotFound = errors.New("session not found")

	ErrSessionAlreadyOpen = errors.New("session already open")

	ErrSessionAlreadyClosed = errors.New("session already closed")

	ErrInvalidClosingCash = errors.New("invalid closing cash")
)
