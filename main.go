package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin engine
	r := gin.Default()

	// Setup base route group
	api := r.Group("/api/v1")
	{
		// Health check endpoint
		api.GET("/health", healthCheck)
	}

	// Start server
	r.Run(":8080")
}

// Health check handler
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "OpenStack Security Hub is running",
	})
}
