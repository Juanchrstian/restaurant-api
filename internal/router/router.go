package router

import (
	"github.com/gin-gonic/gin"

	"github.com/juanchrstian/restaurant-api/internal/health"
)

func New(
	healthHandler *health.Handler,
) *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	{
		api.GET("/health", healthHandler.GetHealth)
	}

	return r
}