package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gunh0/openstack-security-hub/api/handler"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")

	// Health check
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "OpenStack Security Hub is running",
		})
	})

	// Register service-specific routes
	handler.RegisterIdentityRoutes(api)
	// handler.RegisterDashboardRoutes(api)
}
