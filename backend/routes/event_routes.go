package routes

import (
	"backend/handlers"
	
	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(r *gin.Engine) {
	r.POST("/events", handlers.LogEventhandler)
}