package session

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

func (h *Handler) OpenSession(
	c *gin.Context,
) {

	ctx := c.Request.Context()

	var request OpenSessionRequest

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

	session, err := h.service.OpenSession(
		ctx,
		request,
	)

	if err != nil {

		switch err {

		case sharederrors.ErrSessionAlreadyOpen:

			response.Error(
				c,
				http.StatusConflict,
				"SESSION_ALREADY_OPEN",
				"There is already an active session",
			)

		default:

			response.Error(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Failed to open session",
			)

		}

		return

	}

	response.Success(
		c,
		"Session opened successfully",
		ToResponse(session),
	)

}

func (h *Handler) GetActiveSession(
	c *gin.Context,
) {

	ctx := c.Request.Context()

	session, err := h.service.GetActiveSession(ctx)

	if err != nil {

		switch err {

		case sharederrors.ErrSessionNotFound:

			response.Error(
				c,
				http.StatusNotFound,
				"SESSION_NOT_FOUND",
				"No active session",
			)

		default:

			response.Error(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Failed to retrieve session",
			)

		}

		return

	}

	response.Success(
		c,
		"Active session retrieved successfully",
		ToResponse(session),
	)

}

func (h *Handler) CloseSession(
	c *gin.Context,
) {

	ctx := c.Request.Context()

	var request CloseSessionRequest

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

	session, summary, err := h.service.CloseSession(
		ctx,
		request,
	)

	if err != nil {

		switch err {

		case sharederrors.ErrSessionNotFound:

			response.Error(
				c,
				http.StatusNotFound,
				"SESSION_NOT_FOUND",
				"No active session",
			)

		case sharederrors.ErrInvalidClosingCash:

			response.Error(
				c,
				http.StatusBadRequest,
				"INVALID_CLOSING_CASH",
				err.Error(),
			)

		default:

			response.Error(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Failed to close session",
			)

		}

		return
	}

	response.Success(
		c,
		"Session closed successfully",
		ToCloseSessionResponse(
			session,
			summary,
		),
	)
}
