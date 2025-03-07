package main

import (
	"log"
	"santrikoding/backend-api/models"
	"santrikoding/backend-api/routes"

	"github.com/gin-gonic/gin"
)

// @Title Golang App
// @Version 1.0
// @Description This is loreme seipsume.
// @SecurityDefinitions.apikey BearerAuth
// @In header
// @Name Authorization
func main() {
	models.ConnectDatabase()

	app := gin.Default()
	routes.RegisterRoutes(app)

	if err := app.Run(":5000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
