package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateVariantHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	var input struct {
		Name 			string	`json:"name" binding:"required"`
		Weight			float64	`json:"weight" binding:"required"`
		ExperimentID	int		`json:"experiment_id" binding:"required"`
	}

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	var ownerID int
	err := db.DB.Get(&ownerID, "SELECT created_by FROM experiments WHERE id = $1", input.ExperimentID)
	if err != nil || ownerID != userID {
		utils.SendError(c, http.StatusForbidden, "You don't have permission for this resource")
		return
	}

	query := `
		INSERT INTO variants (name, weight, experiment_id, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	var variantID int
	err = db.DB.Get(&variantID, query, input.Name, input.Weight, input.ExperimentID, time.Now())
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create variant")
		return
	}

	utils.SendSuccess(c, http.StatusOK, gin.H{
		"message": "Variant created successfully",
		"variant": gin.H{
			"id": variantID,
			"name": input.Name,
			"weight": input.Weight,
			"experiment_id": input.ExperimentID,
		},
	})
}

func GetVariantsHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	experimentIDStr := c.Query("experiment_id")

	if experimentIDStr == "" {
		utils.SendError(c, http.StatusBadRequest, "experiment_id query param is required")
		return
	}

	experimentID, err := strconv.Atoi(experimentIDStr)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid experiment_id")
		return
	}

	var ownerID int
	err = db.DB.Get(&ownerID, `SELECT created_by FROM experiments WHERE id = $1`, experimentID)
	if err != nil || ownerID != userID {
		utils.SendError(c, http.StatusForbidden, "You don't have permission for this resource")
		return
	}

	var variants []models.Variant
	query := `SELECT id, name, weight, experiment_id, created_at FROM variants WHERE experiment_id = $1 ORDER BY created_at`
	err = db.DB.Select(&variants, query, experimentID)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch variants")
		return
	}

	utils.SendSuccess(c, http.StatusOK, gin.H{"variants": variants})
}

func UpdateVariantHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	variantIDStr := c.Param("id")
	variantID, err := strconv.Atoi(variantIDStr)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid variant ID")
		return
	}

	var input struct {
		Name 	*string	`json:"name"`
		Weight	*float64	`json:"weight"`
	}

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	var ownerID int
	err = db.DB.Get(&ownerID, `
		SELECT e.created_by
		FROM variants v
		JOIN experiments e ON v.experiment_id = e.id
		WHERE v.id = $1
	`, variantID)
	if err != nil || ownerID != userID {
		utils.SendError(c, http.StatusForbidden, "You don't have permission for this resource")
		return
	}

	query := `UPDATE variants SET`
	params := []interface{}{}
	paramIndex := 1

	if input.Name != nil {
		query += `name = $` + strconv.Itoa(paramIndex) + `, `
		params = append(params, *input.Name)
		paramIndex++
	}

	if input.Weight != nil {
		query += `weight = $` + strconv.Itoa(paramIndex) + `, `
		params = append(params, *input.Weight)
		paramIndex++
	}

	if len(params) == 0 {
		utils.SendError(c, http.StatusBadRequest, "No fields provided for update")
		return
	}

	query = query[:len(query)-2]
	query += ` WHERE id = $` + strconv.Itoa(paramIndex)
	params = append(params, variantID)

	_, err = db.DB.Exec(query, params...)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to update variant")
		return
	}

	utils.SendSuccess(c, http.StatusOK, gin.H{"message": "Variant updated successfully"})
}

func DeleteVariantHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	variantIDStr := c.Param("id")
	variantID, err := strconv.Atoi(variantIDStr)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid variant ID")
		return
	}

	var ownerID int
	err = db.DB.Get(&ownerID, `
		SELECT e.created_by
		FROM variants v
		JOIN experiments e ON v.experiment_id = e.id
		WHERE v.id = $1
	`, variantID)

	if err != nil || ownerID != userID {
		utils.SendError(c, http.StatusForbidden, "You don't have permission for this resource")
		return
	}

	_, err = db.DB.Exec(`DELETE FROM variants WHERE id = $1`, variantID)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to delete variant")
		return
	}

	utils.SendSuccess(c, http.StatusOK, gin.H{"message": "Variant deleted successfully"})
}

func AssignVariantHandler(c *gin.Context) {
	var input struct {
		ExperimentID   int    `json:"experiment_id" binding:"required"`
		UserIdentifier string `json:"user_identifier" binding:"required"`
	}

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Try to fetch existing assigned variant
	var existingVariant models.Variant
	err := db.DB.Get(&existingVariant, `
		SELECT v.* FROM variant_assignments va
		JOIN variants v ON va.variant_id = v.id
		WHERE va.user_identifier = $1 AND va.experiment_id = $2
		LIMIT 1
	`, input.UserIdentifier, input.ExperimentID)

	if err == nil {
		utils.SendSuccess(c, http.StatusOK, gin.H{
			"variant": existingVariant,
		})
		return
	}

	if !errors.Is(err, sql.ErrNoRows) {
		utils.SendError(c, http.StatusInternalServerError, "Failed to check existing assignment")
		return
	}

	// Fetch all variants for the experiment
	var variants []models.Variant
	err = db.DB.Select(&variants, `
		SELECT * FROM variants WHERE experiment_id = $1
	`, input.ExperimentID)
	if err != nil || len(variants) == 0 {
		utils.SendError(c, http.StatusNotFound, "No variants found for this experiment")
		return
	}

	// Pick a variant based on weight
	selectedVariantID := utils.PickWeightedVariant(variants)

	// Save assignment
	_, err = db.DB.Exec(`
		INSERT INTO variant_assignments (user_identifier, experiment_id, variant_id)
		VALUES ($1, $2, $3)
	`, input.UserIdentifier, input.ExperimentID, selectedVariantID)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to assign a variant")
		return
	}

	// Fetch full details of assigned variant
	var assignedVariant models.Variant
	err = db.DB.Get(&assignedVariant, `SELECT * FROM variants WHERE id = $1`, selectedVariantID)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch assigned variant")
		return
	}

	utils.SendSuccess(c, http.StatusOK, gin.H{
		"variant": assignedVariant,
	})
}