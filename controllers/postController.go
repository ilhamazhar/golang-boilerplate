package controllers

import (
	"net/http"
	"santrikoding/backend-api/models"
	"santrikoding/backend-api/utils"

	"github.com/gin-gonic/gin"
)

type PostInput struct {
	Title   string `json:"title" binding:"required" form:"title"`
	Content string `json:"content" binding:"required" form:"content"`
}

func CreatePost(c *gin.Context) {
	var input PostInput

	if err := utils.ValidateRequest(c, &input); err != nil {
		return
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}

	if err := models.DB.Create(&post).Error; err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Failed to create post", err)
		return
	}

	utils.JSONResponse(c, http.StatusCreated, "Post created successfully", post)
}

func GetPosts(c *gin.Context) {
	var posts []models.Post

	if err := models.DB.Find(&posts).Error; err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Failed to retrieve posts", err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "List of Posts", posts)
}

func GetPostById(c *gin.Context) {
	var post models.Post

	if err := models.DB.First(&post, c.Param("id")).Error; err != nil {
		utils.JSONError(c, http.StatusNotFound, "Post not found", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Post details", post)
}

func UpdatePostById(c *gin.Context) {
	var post models.Post

	if err := models.DB.First(&post, c.Param("id")).Error; err != nil {
		utils.JSONError(c, http.StatusNotFound, "Post not found", nil)
		return
	}

	var input PostInput
	if err := utils.ValidateRequest(c, &input); err != nil {
		return
	}

	if err := models.DB.Model(&post).Updates(input).Error; err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Failed to update post", err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Post updated successfully", post)
}

func DeletePostById(c *gin.Context) {
	var post models.Post
	if err := models.DB.First(&post, c.Param("id")).Error; err != nil {
		utils.JSONError(c, http.StatusNotFound, "Post not found", nil)
		return
	}

	if err := models.DB.Delete(&post).Error; err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Failed to delete post", err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Post deleted successfully", nil)
}
