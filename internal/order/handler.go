package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sharederrors "github.com/juanchrstian/restaurant-api/internal/shared/errors"
	"github.com/juanchrstian/restaurant-api/internal/shared/response"
	sharedvalidator "github.com/juanchrstian/restaurant-api/internal/shared/validator"
)

type Handler struct {
	service Service
}

func NewHandler(
	service Service,
) *Handler {

	return &Handler{
		service: service,
	}

}

func (h *Handler) CreateOrder(
	c *gin.Context,
) {

	ctx := c.Request.Context()

	order, err := h.service.CreateOrder(ctx)

	if err != nil {

		switch err {

		case sharederrors.ErrSessionNotFound:

			response.Error(
				c,
				http.StatusConflict,
				"NO_ACTIVE_SESSION",
				"Open a session before creating an order",
			)

		default:

			response.Error(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Failed to create order",
			)

		}

		return

	}

	response.Success(
		c,
		"Order created successfully",
		ToResponse(order),
	)

}

func (h *Handler) UpdateItem(
	c *gin.Context,
) {

	ctx := c.Request.Context()

	orderID := c.Param("orderId")
	itemID := c.Param("itemId")

	var request UpdateOrderItemRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			"INVALID_REQUEST",
			"Invalid request body",
		)

		return
	}

	if err := sharedvalidator.Validate(request); err != nil {

		response.Validation(
			c,
			sharedvalidator.ParseErrors(err),
		)

		return
	}

	item, err := h.service.UpdateItem(
		ctx,
		orderID,
		itemID,
		request,
	)

	if err != nil {

		switch err {

		case sharederrors.ErrOrderNotFound:

			response.Error(
				c,
				http.StatusNotFound,
				"ORDER_NOT_FOUND",
				"Order not found",
			)

		case sharederrors.ErrOrderItemNotFound:

			response.Error(
				c,
				http.StatusNotFound,
				"ORDER_ITEM_NOT_FOUND",
				"Order item not found",
			)

		case sharederrors.ErrInsufficientStock:

			response.Error(
				c,
				http.StatusBadRequest,
				"INSUFFICIENT_STOCK",
				"Insufficient stock",
			)

		default:

			response.Error(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Failed to update order item",
			)

		}

		return
	}

	response.Success(
		c,
		"Order item updated successfully",
		ToOrderItemResponse(item),
	)
}

func (h *Handler) AddItem(
	c *gin.Context,
) {

	ctx := c.Request.Context()

	orderID := c.Param("orderId")

	var request AddOrderItemRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			"INVALID_REQUEST",
			"Invalid request body",
		)

		return
	}

	if err := sharedvalidator.Validate(request); err != nil {

		response.Validation(
			c,
			sharedvalidator.ParseErrors(err),
		)

		return
	}

	item, err := h.service.AddItem(
		ctx,
		orderID,
		request,
	)

	if err != nil {

		switch err {

		case sharederrors.ErrMenuUnavailable:

			response.Error(
				c,
				http.StatusBadRequest,
				"MENU_UNAVAILABLE",
				"Menu is unavailable",
			)

		case sharederrors.ErrInsufficientStock:

			response.Error(
				c,
				http.StatusBadRequest,
				"INSUFFICIENT_STOCK",
				"Menu stock is insufficient",
			)

		default:

			response.Error(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Failed to add item",
			)

		}

		return
	}

	response.Success(
		c,
		"Item added successfully",
		ToOrderItemResponse(item),
	)
}

func (h *Handler) GetOrder(
	c *gin.Context,
) {

	ctx := c.Request.Context()

	id := c.Param("orderId")

	order, err := h.service.GetOrder(
		ctx,
		id,
	)

	if err != nil {

		response.Error(
			c,
			http.StatusNotFound,
			"ORDER_NOT_FOUND",
			"Order not found",
		)

		return
	}

	response.Success(
		c,
		"Order retrieved successfully",
		ToDetailResponse(order),
	)
}

func (h *Handler) RemoveItem(
	c *gin.Context,
) {

	ctx := c.Request.Context()

	orderID := c.Param("orderId")
	itemID := c.Param("itemId")

	err := h.service.RemoveItem(
		ctx,
		orderID,
		itemID,
	)
	if err != nil {

		switch err {

		case sharederrors.ErrOrderItemNotFound:

			response.Error(
				c,
				http.StatusNotFound,
				"ORDER_ITEM_NOT_FOUND",
				"Order item not found",
			)

		default:

			response.Error(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Failed to remove item",
			)

		}

		return
	}

	response.Success(
		c,
		"Item removed successfully",
		nil,
	)
}

func (h *Handler) PayOrder(
	c *gin.Context,
) {

	ctx := c.Request.Context()

	orderID := c.Param("orderId")

	var request PaymentRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			"INVALID_REQUEST",
			"Invalid request body",
		)

		return
	}

	if err := sharedvalidator.Validate(request); err != nil {

		response.Validation(
			c,
			sharedvalidator.ParseErrors(err),
		)

		return
	}

	order, err := h.service.PayOrder(
		ctx,
		orderID,
		request,
	)

	if err != nil {

		switch err {

		case sharederrors.ErrOrderAlreadyPaid:

			response.Error(
				c,
				http.StatusConflict,
				"ORDER_ALREADY_PAID",
				"Order already paid",
			)

		case sharederrors.ErrInsufficientPayment:

			response.Error(
				c,
				http.StatusBadRequest,
				"INSUFFICIENT_PAYMENT",
				"Paid amount is less than total amount",
			)

		case sharederrors.ErrEmptyOrder:

			response.Error(
				c,
				http.StatusBadRequest,
				"EMPTY_ORDER",
				"Order has no items",
			)

		case sharederrors.ErrOrderNotFound:

			response.Error(
				c,
				http.StatusNotFound,
				"ORDER_NOT_FOUND",
				"Order not found",
			)

		default:

			response.Error(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Failed to process payment",
			)
		}

		return
	}

	response.Success(
		c,
		"Payment completed successfully",
		ToPaymentResponse(order),
	)
}

func (h *Handler) GetOrders(
	c *gin.Context,
) {

	ctx := c.Request.Context()

	var request GetOrdersRequest

	if err := c.ShouldBindQuery(&request); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			"INVALID_REQUEST",
			"Invalid query parameter",
		)

		return
	}

	if request.Page <= 0 {
		request.Page = 1
	}

	if request.Limit <= 0 {
		request.Limit = 10
	}

	if request.Limit > 100 {
		request.Limit = 100
	}

	if request.Sort == "" {
		request.Sort = "created_at"
	}

	if request.Order == "" {
		request.Order = "desc"
	}

	allowedSort := map[string]bool{
		"created_at":   true,
		"total_amount": true,
		"status":       true,
	}

	if !allowedSort[request.Sort] {
		request.Sort = "created_at"
	}

	switch request.Order {
	case "asc", "ASC":
		request.Order = "ASC"
	default:
		request.Order = "DESC"
	}

	orders, err := h.service.GetAll(
		ctx,
		request,
	)

	if err != nil {

		response.Error(
			c,
			http.StatusInternalServerError,
			"INTERNAL_SERVER_ERROR",
			"Failed to retrieve orders",
		)

		return
	}

	response.Success(
		c,
		"Orders retrieved successfully",
		ToOrderHistoryResponses(orders),
	)
}
