package router

import (
	"github.com/gin-gonic/gin"

	"github.com/juanchrstian/restaurant-api/internal/health"
	"github.com/juanchrstian/restaurant-api/internal/menu"
	"github.com/juanchrstian/restaurant-api/internal/order"
	"github.com/juanchrstian/restaurant-api/internal/session"
)

func New(
	healthHandler *health.Handler,
	menuHandler *menu.Handler,
	sessionHandler *session.Handler,
	orderHandler *order.Handler,
) *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	{

		api.GET("/health", healthHandler.GetHealth)

		menus := api.Group("/menus")
		{
			menus.GET("", menuHandler.GetMenus)
			menus.GET("/:id", menuHandler.GetMenu)
			menus.POST("", menuHandler.CreateMenu)
			menus.PUT("/:id", menuHandler.UpdateMenu)
			menus.DELETE("/:id", menuHandler.DeleteMenu)
		}

		sessions := api.Group("/sessions")
		{
			sessions.POST("", sessionHandler.OpenSession)
			sessions.GET("/active", sessionHandler.GetActiveSession)
			sessions.PATCH("/close", sessionHandler.CloseSession)
		}

		orders := api.Group("/orders")
		{
			orders.POST("", orderHandler.CreateOrder)
			orders.POST("/:id/items", orderHandler.AddItem)
			orders.GET("/:id", orderHandler.GetOrder)
			orders.DELETE("/:orderId/items/:itemId", orderHandler.RemoveItem)
		}

	}

	return router
}
