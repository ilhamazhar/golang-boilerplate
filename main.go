package main

import (
	"log"
	"santrikoding/backend-api/models"
	"santrikoding/backend-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()
	r := gin.Default()
	routes.RegisterRoutes(r)

	if err := r.Run(":5000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
