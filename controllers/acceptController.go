package controllers

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"verve_assignment/models"

	"github.com/gin-gonic/gin"
)

var httpClient = &http.Client{}
var workerPool = make(chan struct{}, 50) // Limit concurrency to 50 workers

// HandlePing handles the ping request for root API
func HandlePing(c *gin.Context) {
	// Get the server's IP address
	ip, err := getServerIP()
	if err != nil {
		// If there's an error, send a response with a failure message
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "unable to get server IP"})
		return
	}

	// Send response with the current server's IP address
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"message":   "pong",
		"server_ip": ip,
	})
}

// getServerIP returns the server's IP address
func getServerIP() (string, error) {
	// Get all network interfaces
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	// Loop through the addresses and find the first non-loopback address
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			// Return the first non-loopback IP address found
			return ipNet.IP.String(), nil
		}
	}

	// If no non-loopback address is found, return an error
	return "", fmt.Errorf("no valid IP address found")
}

func HandleAccept(c *gin.Context) {
	id := c.Query("id")
	endpoint := c.Query("endpoint")

	if id == "" {
		c.String(http.StatusBadRequest, "failed")
		return
	}

	err := models.StoreRequestID(id)
	if err != nil {
		log.Println("Error storing ID in Redis:", err)
		c.String(http.StatusInternalServerError, "failed")
		return
	}

	if endpoint != "" {
		go sendPostRequest(endpoint)
	}

	c.String(http.StatusOK, "ok")
}

// sendPostRequest sends an async POST request without blocking the main request
func sendPostRequest(endpoint string) {
	workerPool <- struct{}{}        // Acquire a worker slot
	defer func() { <-workerPool }() // Release the slot

	// Fire HTTP POST request
	resp, err := httpClient.Post(endpoint, "application/json", nil)
	if err != nil {
		log.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	// Log response status
	log.Printf("POST to %s responded with status: %d\n", endpoint, resp.StatusCode)
}
