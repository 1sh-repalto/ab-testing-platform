package utils

import "github.com/gin-gonic/gin"

func SendSuccess(c *gin.Context, status int, body gin.H) {
	c.JSON(status, gin.H{"success": true, "body": body})
}

func SendError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"success": false, "error": message})
}
