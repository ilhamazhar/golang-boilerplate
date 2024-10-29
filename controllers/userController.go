package controllers

import (
	"net/http"
	"santrikoding/backend-api/models"
	"santrikoding/backend-api/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Get Users
// @Description Get user profile data
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /api/users [get]
func GetUsers(c *gin.Context) {
	var users []models.User

	if err := models.DB.Find(&users).Error; err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Failed to retrieve users", err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "List of Users", users)
}

// @Summary Get user by ID
// @Description Get details of a user by ID
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Param id path int true " "
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Router /api/users/{id} [get]
func GetUserById(c *gin.Context) {
	var user models.User

	if err := models.DB.First(&user, c.Param("id")).Error; err != nil {
		utils.JSONError(c, http.StatusNotFound, "User not found", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "User details", user)
}

// @Summary Update user by ID
// @Description Update user data by ID
// @Tags Users
// @Security BearerAuth
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path int true " "
// @Param name formData string false " "
// @Param email formData string false " "
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Router /api/users/{id} [put]
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

// @Summary Delete user by ID
// @Description Delete user data by ID
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Param id path int true " "
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Router /api/users/{id} [delete]
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
