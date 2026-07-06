package menu

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

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetMenus(c *gin.Context) {

	ctx := c.Request.Context()

	menus, err := h.service.GetMenus(ctx)
	if err != nil {

		response.Error(
			c,
			http.StatusInternalServerError,
			"INTERNAL_SERVER_ERROR",
			"Failed to retrieve menus",
		)

		return
	}

	response.Success(
		c,
		"Menus retrieved successfully",
		ToResponses(menus),
	)
}

func (h *Handler) GetMenu(c *gin.Context) {

	ctx := c.Request.Context()

	id := c.Param("id")

	menu, err := h.service.GetMenu(ctx, id)
	if err != nil {

		switch err {

		case sharederrors.ErrMenuNotFound:

			response.Error(
				c,
				http.StatusNotFound,
				"MENU_NOT_FOUND",
				"Menu not found",
			)

		default:

			response.Error(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Internal server error",
			)

		}

		return
	}

	response.Success(
		c,
		"Menu retrieved successfully",
		ToResponse(menu),
	)
}

func (h *Handler) CreateMenu(c *gin.Context) {

	ctx := c.Request.Context()

	var request CreateMenuRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			"INVALID_REQUEST",
			err.Error(),
		)

		return
	}

	// VALIDATION
	if err := sharedvalidator.Validate(request); err != nil {

		response.Validation(
			c,
			sharedvalidator.ParseErrors(err),
		)

		return
	}

	menu, err := h.service.CreateMenu(ctx, request)
	if err != nil {

		response.Error(
			c,
			http.StatusInternalServerError,
			"INTERNAL_SERVER_ERROR",
			"Failed to create menu",
		)

		return
	}

	response.Success(
		c,
		"Menu created successfully",
		ToResponse(menu),
	)
}

func (h *Handler) UpdateMenu(c *gin.Context) {

	ctx := c.Request.Context()

	id := c.Param("id")

	var request UpdateMenuRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		response.Error(
			c,
			http.StatusBadRequest,
			"INVALID_REQUEST",
			err.Error(),
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

	menu, err := h.service.UpdateMenu(
		ctx,
		id,
		request,
	)

	if err != nil {

		switch err {

		case sharederrors.ErrMenuNotFound:

			response.Error(
				c,
				http.StatusNotFound,
				"MENU_NOT_FOUND",
				"Menu not found",
			)

		default:

			response.Error(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Failed to update menu",
			)

		}

		return
	}

	response.Success(
		c,
		"Menu updated successfully",
		ToResponse(menu),
	)

}
