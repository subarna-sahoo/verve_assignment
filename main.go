package main

import (
	"log"
	"verve_assignment/routes" // Correct import path

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin
	r := gin.Default()

	// Initialize Router
	routes.SetupRoutes(r)

	// Start server
	if err := r.Run(":5000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
