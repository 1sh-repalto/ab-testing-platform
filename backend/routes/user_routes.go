package routes

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	auth := r.Group("/api/auth")
	{
		auth.POST("/signup", handlers.SignupHandler)
		auth.POST("/login", handlers.LoginHandler)
		auth.GET("/me", middleware.AuthMiddleware(), handlers.MeHandler)
		auth.POST("/logout", middleware.AuthMiddleware(), handlers.LogoutHandler)
		auth.POST("/refresh", handlers.RefreshHandler)
	}
}
