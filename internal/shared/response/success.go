package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(
	c *gin.Context,
	message string,
	data interface{},
) {
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}