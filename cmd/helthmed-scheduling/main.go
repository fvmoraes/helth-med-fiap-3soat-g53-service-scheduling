package main

import (
	"helthmed-scheduling/internal/scheduling/infrastructure/db"
	"helthmed-scheduling/internal/scheduling/interfaces/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	db.Init()

	// Set up Gin router
	router := gin.Default()

	// Initialize HTTP handlers and routes
	http.InitRoutes(router)

	// Run the server
	router.Run(":8081")
}
