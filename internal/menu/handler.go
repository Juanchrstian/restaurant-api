package menu

import (
	"log"
	"net/http"
	"strconv"

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

	filter := MenuFilter{
		Search: c.Query("search"),
		SortBy: c.DefaultQuery("sort", "name"),
		Order:  c.DefaultQuery("order", "asc"),
	}

	// Parse available=true|false
	if available := c.Query("available"); available != "" {

		value, err := strconv.ParseBool(available)
		if err == nil {
			filter.Available = &value
		}

	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}

	filter.Page = page
	filter.Limit = limit

	filter.Normalize()

	log.Printf("%+v\n", filter)

	menus, err := h.service.GetMenus(
		ctx,
		filter,
	)
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

func (h *Handler) DeleteMenu(c *gin.Context) {

	ctx := c.Request.Context()

	id := c.Param("id")

	err := h.service.DeleteMenu(ctx, id)

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
				"Failed to delete menu",
			)
		}

		return

	}

	c.Status(http.StatusNoContent)

}
