package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogEventhandler(c *gin.Context) {
	var input models.Event

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if input.EventType != "view" && input.EventType != "conversion" {
		utils.SendError(c, http.StatusBadRequest, "Invalid event type")
		return
	}

	_, err := db.DB.Exec(`
		INSERT INTO events (experiment_id, variant_id, user_identifier, event_type)
		VALUES ($1, $2, $3, $4)
	`, input.ExperimentID, input.VariantID, input.UserIdentifier, input.EventType)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to log event")
		return
	}

	utils.SendSuccess(c, http.StatusOK, gin.H{
		"message": "Event logged",
	})
}
