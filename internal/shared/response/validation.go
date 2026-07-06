package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validation(
	c *gin.Context,
	errors any,
) {

	c.JSON(
		http.StatusBadRequest,
		gin.H{
			"success": false,
			"code":    "VALIDATION_ERROR",
			"message": "Validation failed",
			"errors":  errors,
		},
	)
}
