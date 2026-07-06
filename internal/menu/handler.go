package menu

import (
	"net/http"

	sharederrors "github.com/juanchrstian/restaurant-api/internal/shared/errors"

	"github.com/gin-gonic/gin"
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

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, ToResponses(menus))
}

func (h *Handler) GetMenu(c *gin.Context) {

	ctx := c.Request.Context()

	id := c.Param("id")

	menu, err := h.service.GetMenu(ctx, id)

	if err != nil {

		switch err {

		case sharederrors.ErrMenuNotFound:

			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})

		default:

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}

		return
	}

	c.JSON(http.StatusOK, ToResponse(*menu))
}