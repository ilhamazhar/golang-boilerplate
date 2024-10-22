package controllers

import (
	"net/http"
	"santrikoding/backend-api/models"
	"santrikoding/backend-api/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func Register(c *gin.Context) {
	var input models.User

	if err := utils.ValidateRequest(c, &input); err != nil {
		return
	}

	var existingUser models.User
	if err := models.DB.Where("email = ?", &input.Email).First(&existingUser).Error; err == nil {
		utils.JSONError(c, http.StatusConflict, "Email already registered", nil)
		return
	}

	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Failed to hash password", err)
		return
	}
	input.Password = hashedPassword

	if err := models.DB.Create(&input).Error; err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Failed to register user", err)
		return
	}

	input.Password = ""
	utils.JSONResponse(c, http.StatusCreated, "User registered successfully", input)
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := utils.ValidateRequest(c, &input); err != nil {
		return
	}

	var user models.User
	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		utils.JSONError(c, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		utils.JSONError(c, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Failed to generate token", err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Login successfully", gin.H{"token": token})
}
