package routes

import (
	"santrikoding/backend-api/middleware"
	"santrikoding/backend-api/utils"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		PublicRoutes(api)

		authMiddleware := middleware.AuthMiddleware(utils.GetJwtSecret())
		api.Use(authMiddleware)

		PostRoutes(api)
		UserRoutes(api)
	}
}
