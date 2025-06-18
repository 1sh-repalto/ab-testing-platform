package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var isProd bool = utils.IsProd()
const (
	AccessTokenExpiry  = 60 * 15         	// 15 minutes
	RefreshTokenExpiry = 60 * 60 * 24 * 7 	// 7 days
)

func SignupHandler(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
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

	query := `INSERT INTO users (name, email, password, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	var userID int
	err = db.DB.Get(&userID, query, input.Name, input.Email, string(hashedPassword), time.Now())
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	accessToken, err := utils.GenerateAccessToken(userID)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to generate access token")
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(userID)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to generate refresh token")
		return
	}

	c.SetCookie("access_token", accessToken, AccessTokenExpiry, "/", "", isProd, true)
	c.SetCookie("refresh_token", refreshToken, RefreshTokenExpiry, "/", "", isProd, true)

	utils.SendSuccess(c, http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"id": userID, "name": input.Name,
			"email": input.Email,
		},
	})
}

func LoginHandler(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := models.GetUserByEmail(input.Email)
	if err != nil || user == nil {
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

	accessToken, err := utils.GenerateAccessToken(userID)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(userID)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	c.SetCookie("access_token", accessToken, AccessTokenExpiry, "/", "", isProd, true)
	c.SetCookie("refresh_token", refreshToken, RefreshTokenExpiry, "/", "", isProd, true) // secure=false for dev, change to true in prod

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

	user, err := models.GetUserByID(userID)
	if err != nil || user == nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch user")
		return
	}

	utils.SendSuccess(c, http.StatusOK, gin.H{"user": user})
}

func LogoutHandler(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", isProd, true)
	c.SetCookie("refresh_token", "", -1, "/", "", isProd, true)

	utils.SendSuccess(c, http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func RefreshHandler(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		utils.SendError(c, http.StatusUnauthorized, "No refresh token found")
		return
	}

	_, claims, err := utils.VerifyRefreshToken(refreshToken)
	if err != nil {
		utils.SendError(c, http.StatusUnauthorized, "Invalid or expired refresh token")
		return
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		utils.SendError(c, http.StatusInternalServerError, "Invalid token payload")
		return
	}

	userID := int(userIDFloat)

	newAccessToken, err := utils.GenerateAccessToken(userID)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to generate access token")
		return
	}

	c.SetCookie("access_token", newAccessToken, AccessTokenExpiry, "/", "", isProd, true)

	utils.SendSuccess(c, http.StatusOK, gin.H{"message": "Access token refreshed"})
}
