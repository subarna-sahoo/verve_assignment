package main

import (
	"log"
	"time"
	"verve_assignment/jobs"
	"verve_assignment/routes"
	"verve_assignment/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Redis once
	utils.InitRedis()

	// Initialize RabbitMQ
	utils.InitRabbitMQ()
	defer utils.CloseRabbitMQ()

	go jobs.StartJobScheduler(time.Minute, jobs.UniqueRequestJob)

	// Initialize Gin
	r := gin.Default()

	// Initialize Router
	routes.SetupRoutes(r)

	// Start server
	if err := r.Run(":5000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
