package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"net/http"
	"time"
	"strconv"

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
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to hash password")
		return 
	}

	query := `INSERT INTO users (name, email, password, created_at) VALUES ($1, $2, $3, $4)`
	_, err = db.DB.Exec(query, input.Name, input.Email, string(hashedPassword), time.Now())
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.SendSuccess(c, http.StatusCreated, gin.H{ "message": "User created successfully" })
}

func LoginHandler(c *gin.Context) {
	var input struct {
		Email		string	`json:"email" binding:"required"`
		Password	string 	`json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	query := `SELECT * FROM users WHERE email = $1`
	err := db.DB.Get(&user, query, input.Email)
	if err != nil {
		utils.SendError(c, http.StatusUnauthorized, "Invalid Email or Password")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		utils.SendError(c, http.StatusUnauthorized, "Invalid Email or Password")
		return
	}

	userID, err := strconv.Atoi(user.ID)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Invalid user ID format")
		return
	}

	token, err := utils.GenerateToken(userID)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	c.SetCookie("token", token, 3600*24*7, "/", "", false, true) // secure=false for dev, change to true in prod

	utils.SendSuccess(c, http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func MeHandler(c *gin.Context) {
	userID := c.GetInt("user_id")

	var user models.User
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
	err := db.DB.Get(&user, query, userID)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch user")
		return
	}

	utils.SendSuccess(c, http.StatusOK, gin.H{"user": user})
}

func LogoutHandler(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)

	utils.SendSuccess(c, http.StatusOK, gin.H{"message": "Logged out successfully"})
}