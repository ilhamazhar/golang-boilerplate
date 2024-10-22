package controllers

import (
	"net/http"
	"santrikoding/backend-api/models"
	"santrikoding/backend-api/utils"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	if err := models.DB.Find(&users).Error; err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Failed to retrieve users", err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "List of Users", users)
}

func GetUserById(c *gin.Context) {
	var user models.User

	if err := models.DB.First(&user, c.Param("id")).Error; err != nil {
		utils.JSONError(c, http.StatusNotFound, "User not found", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "User details", user)
}

func UpdateUserById(c *gin.Context) {
	var user models.User

	if err := models.DB.First(&user, c.Param("id")).Error; err != nil {
		utils.JSONError(c, http.StatusNotFound, "User not found", nil)
		return
	}

	if err := utils.ValidateRequest(c, &user); err != nil {
		return
	}

	if err := models.DB.Model(&user).Updates(user).Error; err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Failed to update user", err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "User updated successfully", user)
}

func DeleteUserById(c *gin.Context) {
	var user models.User
	if err := models.DB.First(&user, c.Param("id")).Error; err != nil {
		utils.JSONError(c, http.StatusNotFound, "User not found", nil)
		return
	}

	if err := models.DB.Delete(&user).Error; err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Failed to delete user", err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "User deleted successfully", nil)
}
