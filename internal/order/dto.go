package order

import "time"

type CreateOrderRequest struct {
}

type OrderResponse struct {
	ID string `json:"id"`

	OrderNumber string `json:"order_number"`

	SessionID string `json:"session_id"`

	Status string `json:"status"`

	TotalAmount int64 `json:"total_amount"`

	CreatedAt time.Time `json:"created_at"`
}

type AddOrderItemRequest struct {
	MenuID   string `json:"menu_id" validate:"required,uuid"`
	Quantity int    `json:"quantity" validate:"required,min=1"`
}

type OrderItemResponse struct {
	ID       string `json:"id"`
	MenuID   string `json:"menu_id"`
	Quantity int    `json:"quantity"`
	Price    int64  `json:"price"`
	Subtotal int64  `json:"subtotal"`
}

func ToResponse(
	order *Order,
) OrderResponse {

	return OrderResponse{

		ID: order.ID.String(),

		OrderNumber: order.OrderNumber,

		SessionID: order.SessionID.String(),

		Status: string(order.Status),

		TotalAmount: order.TotalAmount,

		CreatedAt: order.CreatedAt,
	}
}

func ToOrderItemResponse(
	item *OrderItem,
) OrderItemResponse {

	return OrderItemResponse{
		ID: item.ID.String(),

		MenuID: item.MenuID.String(),

		Quantity: item.Quantity,

		Price: item.Price,

		Subtotal: item.Subtotal,
	}
}

type OrderDetailResponse struct {
	ID string `json:"id"`

	OrderNumber string `json:"order_number"`

	Status string `json:"status"`

	TotalAmount int64 `json:"total_amount"`

	Items []OrderItemDetailResponse `json:"items"`
}

type OrderItemDetailResponse struct {
	ID string `json:"id"`

	MenuID string `json:"menu_id"`

	MenuName string `json:"menu_name"`

	Quantity int `json:"quantity"`

	Price int64 `json:"price"`

	Subtotal int64 `json:"subtotal"`
}

func ToDetailResponse(
	order *Order,
) OrderDetailResponse {

	items := make(
		[]OrderItemDetailResponse,
		0,
		len(order.Items),
	)

	for _, item := range order.Items {

		items = append(
			items,
			OrderItemDetailResponse{
				ID: item.ID.String(),

				MenuID: item.MenuID.String(),

				MenuName: item.Menu.Name,

				Quantity: item.Quantity,

				Price: item.Price,

				Subtotal: item.Subtotal,
			},
		)

	}

	return OrderDetailResponse{

		ID: order.ID.String(),

		OrderNumber: order.OrderNumber,

		Status: string(order.Status),

		TotalAmount: order.TotalAmount,

		Items: items,
	}
}

type UpdateOrderItemRequest struct {
	Quantity int `json:"quantity" validate:"required,min=1"`
}

type PaymentRequest struct {
	PaymentMethod PaymentMethod `json:"payment_method" validate:"required"`
	PaidAmount    int64         `json:"paid_amount" validate:"required,min=1"`
}

type PaymentResponse struct {
	OrderID string `json:"order_id"`

	PaymentMethod string `json:"payment_method"`

	TotalAmount int64 `json:"total_amount"`

	PaidAmount int64 `json:"paid_amount"`

	ChangeAmount int64 `json:"change_amount"`

	Status string `json:"status"`
}

func ToPaymentResponse(
	order *Order,
) PaymentResponse {

	return PaymentResponse{
		OrderID: order.ID.String(),

		PaymentMethod: string(*order.PaymentMethod),

		TotalAmount: order.TotalAmount,

		PaidAmount: *order.PaidAmount,

		ChangeAmount: *order.ChangeAmount,

		Status: string(order.Status),
	}
}

type OrderHistoryResponse struct {
	ID            string    `json:"id"`
	SessionID     string    `json:"session_id"`
	Status        string    `json:"status"`
	PaymentMethod *string   `json:"payment_method,omitempty"`
	TotalAmount   int64     `json:"total_amount"`
	CreatedAt     time.Time `json:"created_at"`
}

func ToOrderHistoryResponse(
	order Order,
) OrderHistoryResponse {

	var paymentMethod *string

	if order.PaymentMethod != nil {
		pm := string(*order.PaymentMethod)
		paymentMethod = &pm
	}

	return OrderHistoryResponse{
		ID:            order.ID.String(),
		SessionID:     order.SessionID.String(),
		Status:        string(order.Status),
		PaymentMethod: paymentMethod,
		TotalAmount:   order.TotalAmount,
		CreatedAt:     order.CreatedAt,
	}
}

func ToOrderHistoryResponses(
	orders []Order,
) []OrderHistoryResponse {

	responses := make([]OrderHistoryResponse, 0, len(orders))

	for _, order := range orders {
		responses = append(
			responses,
			ToOrderHistoryResponse(order),
		)
	}

	return responses
}

type GetOrdersRequest struct {
	Page          int    `form:"page"`
	Limit         int    `form:"limit"`
	Status        string `form:"status"`
	PaymentMethod string `form:"payment_method"`
	SessionID     string `form:"session_id"`
	Sort          string `form:"sort"`
	Order         string `form:"order"`
}
