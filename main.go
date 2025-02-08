package main

import (
	"log"

	"awesomeProject/config"
	"awesomeProject/handlers"
	"awesomeProject/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize dependencies
	config.Init()
	defer config.Cleanup()

	// Start logging unique request counts in a separate goroutine
	go services.LogUniqueCounts()

	// Set up Gin router
	r := gin.Default()
	r.GET("/api/verve/accept", handlers.AcceptHandler)
	r.POST("/get", handlers.GetHandler)

	// Run the server
	log.Println("Server running on port 8080")
	r.Run(":8080")
}
