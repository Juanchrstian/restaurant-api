package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sharederrors "github.com/juanchrstian/restaurant-api/internal/shared/errors"
	"github.com/juanchrstian/restaurant-api/internal/shared/response"
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
				"No active session",
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
