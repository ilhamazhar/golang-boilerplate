package routes

import (
	"santrikoding/backend-api/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(rg *gin.RouterGroup) {
	public := rg.Group("/public")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}
}
