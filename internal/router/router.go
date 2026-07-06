package router

import (
	"github.com/gin-gonic/gin"

	"github.com/juanchrstian/restaurant-api/internal/health"
	"github.com/juanchrstian/restaurant-api/internal/menu"
)

func New(
	healthHandler *health.Handler,
	menuHandler *menu.Handler,
) *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	{
		api.GET("/health", healthHandler.GetHealth)

		api.GET("/menus", menuHandler.GetMenus)
		api.GET("/menus/:id", menuHandler.GetMenu)
	}

	return router
}