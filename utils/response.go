package utils

import (
	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func JSONError(c *gin.Context, status int, message string, errors interface{}) {
	c.AbortWithStatusJSON(status, gin.H{
		"success": false,
		"message": message,
		"errors":  errors,
	})
}
