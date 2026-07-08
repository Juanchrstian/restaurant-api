package session

import (
	"context"
	"time"

	sharederrors "github.com/juanchrstian/restaurant-api/internal/shared/errors"
)

var _ Service = (*service)(nil)

type service struct {
	repository Repository
}

func NewService(
	repository Repository,
) Service {

	return &service{
		repository: repository,
	}

}

// =========================================
// OPEN SESSION
// =========================================

func (s *service) OpenSession(
	ctx context.Context,
	request OpenSessionRequest,
) (*Session, error) {

	active, err := s.repository.GetActive(ctx)

	if err != nil {

		// selain "session tidak ditemukan"
		// berarti benar-benar terjadi error database
		if err != sharederrors.ErrSessionNotFound {
			return nil, err
		}

	} else if active != nil {

		// masih ada session aktif
		return nil, sharederrors.ErrSessionAlreadyOpen

	}

	session := &Session{
		Status:      StatusOpen,
		OpenedBy:    request.OpenedBy,
		OpeningCash: request.OpeningCash,
		OpenedAt:    time.Now(),
	}

	if err := s.repository.Create(
		ctx,
		session,
	); err != nil {

		return nil, err

	}

	return session, nil
}

// =========================================
// GET ACTIVE SESSION
// =========================================

func (s *service) GetActiveSession(
	ctx context.Context,
) (*Session, error) {

	return s.repository.GetActive(ctx)

}

// =========================================
// CLOSE SESSION
// =========================================

func (s *service) CloseSession(
	ctx context.Context,
	request CloseSessionRequest,
) (*Session, error) {

	// Business Validation
	if request.ClosingCash < 0 {
		return nil, sharederrors.ErrInvalidClosingCash
	}

	session, err := s.repository.GetActive(ctx)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	session.Status = StatusClosed
	session.ClosingCash = &request.ClosingCash
	session.ClosedAt = &now

	if err := s.repository.Update(
		ctx,
		session,
	); err != nil {

		return nil, err

	}

	return session, nil
}
