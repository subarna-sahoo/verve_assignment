package services

import (
	"log"
	"net/http"
	"strconv" // Required for converting int to string
	"sync"
	"verve_assignment/models"
)

var uniqueRequests sync.Map

// ProcessRequest processes the incoming request by checking the uniqueness of the ID and sending the count to the endpoint.
func ProcessRequest(id string, endpoint string) string {
	// Check if the ID is unique using Redis (Check if ID is in Redis)
	isUnique, err := models.CheckUniqueID(id)
	if err != nil {
		log.Println("Redis error:", err)
		return "failed"
	}

	// If unique, store it in the local memory map
	if isUnique {
		uniqueRequests.Store(id, struct{}{})
	}

	// Send the unique count to the provided endpoint if the endpoint is not empty
	if endpoint != "" {
		go sendCountToEndpoint(endpoint)
	}

	return "ok"
}

// sendCountToEndpoint sends the count of unique requests to the provided endpoint.
func sendCountToEndpoint(endpoint string) {
	// Count unique requests in memory
	count := 0
	uniqueRequests.Range(func(_, _ interface{}) bool {
		count++
		return true
	})

	// Convert count to string correctly using strconv.Itoa
	countStr := strconv.Itoa(count)

	// Send the count to the endpoint as a query parameter
	resp, err := http.Post(endpoint+"?count="+countStr, "application/json", nil)
	if err != nil {
		log.Println("Failed to send count:", err)
	} else {
		log.Println("Sent count, status code:", resp.StatusCode)
	}
}
