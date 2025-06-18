package handlers

import (
	"backend/db"
	"backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(c *gin.Context) {
	var input struct {
		Name 		string	`json:"name" binding:"required"`
		Email		string	`json:"email" binding:"required,email"`
		Password	string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to hash password")
		return 
	}

	query := `INSERT INTO users (name, email, password, created_at) VALUES ($1, $2, $3, $4)`
	_, err = db.DB.Exec(query, input.Name, input.Email, string(hashedPassword), time.Now())
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.Success(c, http.StatusCreated, "User created successfully")
}