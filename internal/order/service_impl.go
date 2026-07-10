package order

import (
	"context"
	"time"

	"github.com/juanchrstian/restaurant-api/internal/menu"
	"github.com/juanchrstian/restaurant-api/internal/session"
	sharederrors "github.com/juanchrstian/restaurant-api/internal/shared/errors"
	"gorm.io/gorm"
)

var _ Service = (*service)(nil)

type service struct {
	db *gorm.DB

	repository Repository

	menuRepository menu.Repository

	sessionService session.Service
}

func NewService(
	db *gorm.DB,
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

	order, err := s.repository.GetByID(ctx, orderID)
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
	// CHECK EXISTING ITEM
	// =====================================

	item, err := s.repository.GetItemByMenu(
		ctx,
		orderID,
		request.MenuID,
	)

	if err != nil {
		return nil, err
	}

	// =====================================
	// ITEM ALREADY EXISTS
	// =====================================

	if item != nil {

		item.Quantity += request.Quantity

		item.Subtotal = int64(item.Quantity) * item.Price

		if err := s.repository.UpdateItem(
			ctx,
			item,
		); err != nil {

			return nil, err
		}

		order.TotalAmount += int64(request.Quantity) * menu.Price

		if err := s.repository.Update(
			ctx,
			order,
		); err != nil {

			return nil, err
		}

		return item, nil
	}

	// =====================================
	// CREATE NEW ITEM
	// =====================================

	subtotal := int64(request.Quantity) * menu.Price

	item = &OrderItem{
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

func (s *service) UpdateItem(
	ctx context.Context,
	orderID string,
	itemID string,
	request UpdateOrderItemRequest,
) (*OrderItem, error) {

	order, err := s.repository.GetByID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	item, err := s.repository.GetItemByID(ctx, itemID)
	if err != nil {
		return nil, err
	}

	if item.OrderID != order.ID {
		return nil, sharederrors.ErrOrderItemNotFound
	}

	menu, err := s.menuRepository.GetByID(
		ctx,
		item.MenuID.String(),
	)
	if err != nil {
		return nil, err
	}

	if menu.Stock < request.Quantity {
		return nil, sharederrors.ErrInsufficientStock
	}

	oldSubtotal := item.Subtotal
	newSubtotal := int64(request.Quantity) * item.Price

	item.Quantity = request.Quantity
	item.Subtotal = newSubtotal

	order.TotalAmount =
		order.TotalAmount -
			oldSubtotal +
			newSubtotal

	if err := s.repository.UpdateItem(
		ctx,
		item,
	); err != nil {

		return nil, err
	}

	if err := s.repository.Update(
		ctx,
		order,
	); err != nil {

		return nil, err
	}

	return item, nil
}

func (s *service) RemoveItem(
	ctx context.Context,
	orderID string,
	itemID string,
) error {

	order, err := s.repository.GetByID(ctx, orderID)
	if err != nil {
		return err
	}

	item, err := s.repository.GetItemByID(ctx, itemID)
	if err != nil {
		return err
	}

	if item.OrderID != order.ID {
		return sharederrors.ErrOrderItemNotFound
	}

	order.TotalAmount -= item.Subtotal

	if order.TotalAmount < 0 {
		order.TotalAmount = 0
	}

	if err := s.repository.DeleteItem(ctx, item); err != nil {
		return err
	}

	if err := s.repository.Update(ctx, order); err != nil {
		return err
	}

	return nil
}

func (s *service) PayOrder(
	ctx context.Context,
	orderID string,
	request PaymentRequest,
) (*Order, error) {

	var paidOrder *Order

	err := s.db.Transaction(func(tx *gorm.DB) error {

		// =====================================
		// TRANSACTION REPOSITORIES
		// =====================================

		orderRepository := s.repository.WithTransaction(tx)
		menuRepository := s.menuRepository.WithTransaction(tx)

		// =====================================
		// GET ORDER
		// =====================================

		order, err := orderRepository.GetByID(
			ctx,
			orderID,
		)
		if err != nil {
			return err
		}

		// =====================================
		// VALIDATION
		// =====================================

		if order.Status == OrderPaid {
			return sharederrors.ErrOrderAlreadyPaid
		}

		if order.TotalAmount == 0 {
			return sharederrors.ErrEmptyOrder
		}

		if request.PaidAmount < order.TotalAmount {
			return sharederrors.ErrInsufficientPayment
		}

		// =====================================
		// PAYMENT
		// =====================================

		change := request.PaidAmount - order.TotalAmount
		now := time.Now()

		order.PaymentMethod = &request.PaymentMethod
		order.PaidAmount = &request.PaidAmount
		order.ChangeAmount = &change
		order.PaidAt = &now
		order.Status = OrderPaid

		// =====================================
		// UPDATE ORDER
		// =====================================

		if err := orderRepository.Update(
			ctx,
			order,
		); err != nil {
			return err
		}

		// =====================================
		// GET ORDER ITEMS
		// =====================================

		items, err := orderRepository.GetItems(
			ctx,
			order.ID.String(),
		)
		if err != nil {
			return err
		}

		// =====================================
		// REDUCE MENU STOCK
		// =====================================

		for _, item := range items {

			menu, err := menuRepository.GetByID(
				ctx,
				item.MenuID.String(),
			)
			if err != nil {
				return err
			}

			if menu.Stock < item.Quantity {
				return sharederrors.ErrInsufficientStock
			}

			menu.Stock -= item.Quantity

			if err := menuRepository.Update(
				ctx,
				menu,
			); err != nil {
				return err
			}
		}

		paidOrder = order

		return nil
	})

	if err != nil {
		return nil, err
	}

	return paidOrder, nil
}
