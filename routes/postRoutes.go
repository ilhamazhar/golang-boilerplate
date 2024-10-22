package routes

import (
	"santrikoding/backend-api/controllers"

	"github.com/gin-gonic/gin"
)

func PostRoutes(rg *gin.RouterGroup) {
	posts := rg.Group("/posts")
	{
		posts.POST("", controllers.CreatePost)
		posts.GET("", controllers.GetPosts)
		posts.GET("/:id", controllers.GetPostById)
		posts.PUT("/:id", controllers.UpdatePostById)
		posts.DELETE("/:id", controllers.DeletePostById)
	}
}
