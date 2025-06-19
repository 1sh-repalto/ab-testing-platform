package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateExperimentHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	var input struct {
		Name 		string `json:"name" binding:"required"`
		Description	string `json:"description"`
	}

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	query := `
		INSERT INTO experiments (name, description, status, created_by, created_at)
		VALUES ($1, $2, 'draft', $3, $4)
		RETURNING id
	`

	var experimentID int
	err := db.DB.Get(&experimentID, query, input.Name, input.Description, userID, time.Now())
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create experiment")
		return
	}

	utils.SendSuccess(c, http.StatusOK, gin.H{
		"message": "Experiment created",
		"experiment": gin.H{
			"id":          experimentID,
			"name":        input.Name,
			"description": input.Description,
			"status":      "draft",
			"created_by":  userID,
		},
	})
}

func GetUserExperimentsHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	query := `
		SELECT id, name, description, status, created_by, created_at
		FROM experiments
		WHERE created_by = $1
		ORDER BY created_at DESC
	`

	var experiments []models.Experiment
	err := db.DB.Select(&experiments, query, userID)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch experiments")
		return
	}

	utils.SendSuccess(c, http.StatusOK, gin.H { "experiments": experiments })
}

func UpdateExperimentHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	experimentID := c.Param("id")

	var input struct {
		Name        *string `json:"name"`
		Description *string `json:"description"`
		Status      *string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	validStatuses := map[string]bool{"draft": true, "running": true, "paused": true, "completed": true}
	if input.Status != nil && !validStatuses[*input.Status] {
		utils.SendError(c, http.StatusBadRequest, "Invalid status value")
		return
	}

	var queryBuilder strings.Builder
	queryBuilder.WriteString("UPDATE experiments SET ")

	params := []interface{}{}
	paramIndex := 1

	if input.Name != nil {
		queryBuilder.WriteString("name = $" + strconv.Itoa(paramIndex) + ", ")
		params = append(params, *input.Name)
		paramIndex++
	}
	if input.Description != nil {
		queryBuilder.WriteString("description = $" + strconv.Itoa(paramIndex) + ", ")
		params = append(params, *input.Description)
		paramIndex++
	}
	if input.Status != nil {
		queryBuilder.WriteString("status = $" + strconv.Itoa(paramIndex) + ", ")
		params = append(params, *input.Status)
		paramIndex++
	}

	if len(params) == 0 {
		utils.SendError(c, http.StatusBadRequest, "No fields provided for update")
		return
	}

	// Remove the final comma + space
	query := strings.TrimSuffix(queryBuilder.String(), ", ")

	// Convert experimentID to int
	id, err := strconv.Atoi(experimentID)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid experiment ID")
		return
	}

	// Add WHERE clause
	query += " WHERE id = $" + strconv.Itoa(paramIndex) + " AND created_by = $" + strconv.Itoa(paramIndex+1)
	params = append(params, id, userID)

	result, err := db.DB.Exec(query, params...)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to update experiment: "+err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to confirm update")
		return
	}

	if rowsAffected == 0 {
		utils.SendError(c, http.StatusNotFound, "Experiment not found or you are not authorized to update it")
		return
	}

	utils.SendSuccess(c, http.StatusOK, gin.H{"message": "Experiment updated successfully"})
}

func DeleteExperimentHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	experimentID := c.Param("id")

	query := `
		DELETE FROM experiments
		WHERE id = $1 AND created_by = $2
	`

	result, err := db.DB.Exec(query, experimentID, userID)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to delete experiment")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to confirm deletion")
		return
	}

	if rowsAffected == 0 {
		utils.SendError(c, http.StatusNotFound, "Experiment not found or not authorized")
		return
	}

	utils.SendSuccess(c, http.StatusOK, gin.H{"message": "Experiment deleted successfully"})
}