package utils

import "github.com/gin-gonic/gin"

func SendSuccess(c *gin.Context, status int, data gin.H) {
	c.JSON(status, gin.H{"success": true, "data": data})
}

func SendError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"success": false, "error": message})
}