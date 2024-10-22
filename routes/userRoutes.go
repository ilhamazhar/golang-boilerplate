package routes

import (
	"santrikoding/backend-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.GET("", controllers.GetUsers)
		users.GET("/:id", controllers.GetUserById)
		users.PUT("/:id", controllers.UpdateUserById)
		users.DELETE("/:id", controllers.DeleteUserById)
	}
}
