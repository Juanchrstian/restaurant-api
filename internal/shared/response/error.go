package response

import (
	"github.com/gin-gonic/gin"
)

func Error(
	c *gin.Context,
	status int,
	code string,
	message string,
) {
	c.JSON(status, ErrorResponse{
		Success: false,
		Code:    code,
		Message: message,
	})
}