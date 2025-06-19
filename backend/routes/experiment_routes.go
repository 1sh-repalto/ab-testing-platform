package routes

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterExperimentRoutes(r *gin.Engine) {
	experiment := r.Group("/experiments", middleware.AuthMiddleware())
	{
		experiment.POST("/", handlers.CreateExperimentHandler)
		experiment.GET("/", handlers.GetUserExperimentsHandler)
		experiment.PUT("/:id", handlers.UpdateExperimentHandler)
		experiment.DELETE("/:id", handlers.DeleteExperimentHandler)
	}
}