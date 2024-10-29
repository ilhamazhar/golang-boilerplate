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

// @Summary Create Post
// @Description Create a new post
// @Tags Posts
// @Security BearerAuth
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param title formData string true " "
// @Param content formData string true " "
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/posts [post]
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

// @Summary Get Posts
// @Description Get post profile data
// @Tags Posts
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /api/posts [get]
func GetPosts(c *gin.Context) {
	var posts []models.Post

	if err := models.DB.Find(&posts).Error; err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "Failed to retrieve posts", err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "List of Posts", posts)
}

// @Summary Get post by ID
// @Description Get details of a post by ID
// @Tags Posts
// @Security BearerAuth
// @Produce json
// @Param id path int true " "
// @Success 200 {object} models.Post
// @Failure 404 {object} map[string]string
// @Router /api/posts/{id} [get]
func GetPostById(c *gin.Context) {
	var post models.Post

	if err := models.DB.First(&post, c.Param("id")).Error; err != nil {
		utils.JSONError(c, http.StatusNotFound, "Post not found", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Post details", post)
}

// @Summary Update post by ID
// @Description Update post data by ID
// @Tags Posts
// @Security BearerAuth
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path int true " "
// @Param title formData string false " "
// @Param content formData string false " "
// @Success 200 {object} models.Post
// @Failure 404 {object} map[string]string
// @Router /api/posts/{id} [put]
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

// @Summary Delete post by ID
// @Description Delete post data by ID
// @Tags Posts
// @Security BearerAuth
// @Produce json
// @Param id path int true " "
// @Success 200 {object} models.Post
// @Failure 404 {object} map[string]string
// @Router /api/posts/{id} [delete]
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
