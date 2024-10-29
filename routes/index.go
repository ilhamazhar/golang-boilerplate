package routes

import (
	"santrikoding/backend-api/middleware"
	"santrikoding/backend-api/utils"

	_ "santrikoding/backend-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	swaggerUrl := ginSwagger.URL("http://localhost:5000/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerUrl))

	api := router.Group("/api")
	{
		PublicRoutes(api)

		authMiddleware := middleware.AuthMiddleware(utils.GetJwtSecret())
		api.Use(authMiddleware)

		PostRoutes(api)
		UserRoutes(api)
	}
}
