package routes

import (
	"verve_assignment/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Define the GET endpoint
	r.GET("/", controllers.HandlePing)
	r.GET("/api/verve/accept", controllers.HandleAccept)
	r.GET("/api/verve/unique-requests", controllers.GetUniqueRequests) // Just to Test the UniqueRequests counts
}
