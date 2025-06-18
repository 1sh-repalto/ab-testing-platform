package routes

import (
	"github.com/gin-gonic/gin"
	"backend/handlers"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/signup", handlers.SignupHandler)
}