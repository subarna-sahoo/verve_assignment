package controllers

import (
	"log"
	"net/http"

	"verve_assignment/models"

	"github.com/gin-gonic/gin"
)

// GetUniqueRequests handles the /api/verve/unique-requests endpoint
func GetUniqueRequests(c *gin.Context) {
	ids, err := models.GetUniqueRequestIDsFromLastMinute()
	if err != nil {
		log.Println("Error fetching unique requests:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(ids),
		"ids":   ids,
	})
}
