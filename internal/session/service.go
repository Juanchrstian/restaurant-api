package session

import "context"

type Service interface {
	OpenSession(
		ctx context.Context,
		request OpenSessionRequest,
	) (*Session, error)

	GetActiveSession(
		ctx context.Context,
	) (*Session, error)

	CloseSession(
		ctx context.Context,
		request CloseSessionRequest,
	) (*Session, error)
}
