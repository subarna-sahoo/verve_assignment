package controllers

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	fmt.Println(id, endpoint)
	// if id == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": "id is required"})
	// 	return
	// }

	// // Process the request through service
	// status := services.ProcessRequest(id, endpoint)

	// // Send response based on the status
	// if status == "ok" {
	// 	c.String(http.StatusOK, "ok")
	// } else {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": "server error"})
	// }

	type Response struct {
		Status  string
		Message string
		Data    interface{}
	}

	// If ID is missing, return an error response
	if id == "" {
		response := Response{
			Status:  "failed",
			Message: "id is required",
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := Response{
		Status:  "ok",
		Message: "success",
		Data:    id, // Send the ID as part of the response
	}

	// Return the response with status OK
	c.JSON(http.StatusOK, response)
}
