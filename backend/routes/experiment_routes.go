package routes

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterExperimentRoutes(r *gin.Engine) {
	experiment := r.Group("/api", middleware.AuthMiddleware())
	{
		experiment.POST("/experiments", handlers.CreateExperimentHandler)
		experiment.GET("/experiments", handlers.GetUserExperimentsHandler)
		experiment.PUT("/:id", handlers.UpdateExperimentHandler)
		experiment.DELETE("/:id", handlers.DeleteExperimentHandler)
	}
}