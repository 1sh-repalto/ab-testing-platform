package routes

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/signup", handlers.SignupHandler)
	r.POST("/login", handlers.LoginHandler)
	r.GET("/me", middleware.AuthMiddleware(), handlers.MeHandler)
	r.POST("/logout", middleware.AuthMiddleware(), handlers.LogoutHandler)
}