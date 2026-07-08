package session

import "context"

type Repository interface {
	Create(
		ctx context.Context,
		session *Session,
	) error

	GetActive(
		ctx context.Context,
	) (*Session, error)

	Update(
		ctx context.Context,
		session *Session,
	) error
}
