package routes

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterVariantRoutes(r *gin.Engine) {
	variants := r.Group("/variants", middleware.AuthMiddleware())
	{
		variants.POST("/", handlers.CreateVariantHandler)
		variants.GET("/", handlers.GetVariantsHandler)
		variants.PUT("/:id", handlers.UpdateVariantHandler)
		variants.DELETE("/:id", handlers.DeleteVariantHandler)
		variants.POST("/assign", handlers.AssignVariantHandler)
	}
}