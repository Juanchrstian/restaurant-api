package order

import (
	"context"

	"github.com/juanchrstian/restaurant-api/internal/session"
)

var _ Service = (*service)(nil)

type service struct {
	repository Repository

	sessionService session.Service
}

func NewService(
	repository Repository,
	sessionService session.Service,
) Service {

	return &service{
		repository:     repository,
		sessionService: sessionService,
	}

}

func (s *service) CreateOrder(
	ctx context.Context,
) (*Order, error) {

	activeSession, err := s.sessionService.GetActiveSession(ctx)
	if err != nil {
		return nil, err
	}

	order := &Order{
		SessionID:   activeSession.ID,
		OrderNumber: GenerateOrderNumber(),
		Status:      OrderPending,
		TotalAmount: 0,
	}

	if err := s.repository.Create(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}
