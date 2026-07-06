package menu

import (
	"net/http"

	"github.com/gin-gonic/gin"

	sharederrors "github.com/juanchrstian/restaurant-api/internal/shared/errors"
	"github.com/juanchrstian/restaurant-api/internal/shared/response"
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