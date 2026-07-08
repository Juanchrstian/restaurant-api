package order

type CreateOrderRequest struct {
}

type OrderResponse struct {
	ID string `json:"id"`

	OrderNumber string `json:"order_number"`

	Status string `json:"status"`

	TotalAmount int64 `json:"total_amount"`
}

func ToResponse(
	order *Order,
) OrderResponse {

	return OrderResponse{

		ID: order.ID.String(),

		OrderNumber: order.OrderNumber,

		Status: string(order.Status),

		TotalAmount: order.TotalAmount,
	}

}
