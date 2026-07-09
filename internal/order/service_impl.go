package order

import (
	"context"

	"github.com/juanchrstian/restaurant-api/internal/menu"
	"github.com/juanchrstian/restaurant-api/internal/session"
	sharederrors "github.com/juanchrstian/restaurant-api/internal/shared/errors"
)

var _ Service = (*service)(nil)

type service struct {
	repository Repository

	menuRepository menu.Repository

	sessionService session.Service
}

func NewService(
	repository Repository,
	menuRepository menu.Repository,
	sessionService session.Service,
) Service {

	return &service{
		repository:     repository,
		menuRepository: menuRepository,
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

func (s *service) AddItem(
	ctx context.Context,
	orderID string,
	request AddOrderItemRequest,
) (*OrderItem, error) {

	// =====================================
	// GET ORDER
	// =====================================

	order, err := s.repository.GetByID(
		ctx,
		orderID,
	)

	if err != nil {
		return nil, err
	}

	// =====================================
	// GET MENU
	// =====================================

	menu, err := s.menuRepository.GetByID(
		ctx,
		request.MenuID,
	)

	if err != nil {
		return nil, err
	}

	// =====================================
	// VALIDATION
	// =====================================

	if !menu.Available {
		return nil, sharederrors.ErrMenuUnavailable
	}

	if menu.Stock < request.Quantity {
		return nil, sharederrors.ErrInsufficientStock
	}

	// =====================================
	// CALCULATE SUBTOTAL
	// =====================================

	subtotal := int64(request.Quantity) * menu.Price

	// =====================================
	// CREATE ORDER ITEM
	// =====================================

	item := &OrderItem{
		OrderID:  order.ID,
		MenuID:   menu.ID,
		Quantity: request.Quantity,
		Price:    menu.Price,
		Subtotal: subtotal,
	}

	if err := s.repository.CreateItem(
		ctx,
		item,
	); err != nil {

		return nil, err
	}

	// =====================================
	// UPDATE ORDER TOTAL
	// =====================================

	order.TotalAmount += subtotal

	if err := s.repository.Update(
		ctx,
		order,
	); err != nil {

		return nil, err
	}

	return item, nil
}

func (s *service) GetOrder(
	ctx context.Context,
	id string,
) (*Order, error) {

	return s.repository.GetDetail(
		ctx,
		id,
	)

}
