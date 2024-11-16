package identity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
)

// RegisterRoutes registers all identity check routes
func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/check/identity-01", handleIdentity01)
	router.GET("/check/identity-01-01", checkIdentity0101)
	router.GET("/check/identity-01-02", checkIdentity0102)
	router.GET("/check/identity-01-03", checkIdentity0103)
	router.GET("/check/identity-01-04", checkIdentity0104)
	router.GET("/check/identity-01-05", checkIdentity0105)
	router.GET("/check/identity-01-06", checkIdentity0106)
	router.GET("/check/identity-01-07", checkIdentity0107)
	router.GET("/check/identity-01-08", checkIdentity0108)
}

func handleIdentity01(c *gin.Context) {
	client, err := getSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer client.Close()

	checks := []struct {
		name string
		fn   func(*ssh.Client) CheckResult
	}{
		{"Identity-01-01", CheckIdentity0101},
		{"Identity-01-02", CheckIdentity0102},
		{"Identity-01-03", CheckIdentity0103},
		{"Identity-01-04", CheckIdentity0104},
		{"Identity-01-05", CheckIdentity0105},
		{"Identity-01-06", CheckIdentity0106},
		{"Identity-01-07", CheckIdentity0107},
		{"Identity-01-08", CheckIdentity0108},
	}

	var results []map[string]CheckResult
	for _, check := range checks {
		result := check.fn(client)
		results = append(results, map[string]CheckResult{check.name: result})
	}

	c.JSON(http.StatusOK, results)
}

func checkIdentity0101(c *gin.Context) {
	client, err := getSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := CheckIdentity0101(client)
	c.JSON(http.StatusOK, result)
}

func checkIdentity0102(c *gin.Context) {
	client, err := getSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := CheckIdentity0102(client)
	c.JSON(http.StatusOK, result)
}

func checkIdentity0103(c *gin.Context) {
	client, err := getSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := CheckIdentity0103(client)
	c.JSON(http.StatusOK, result)
}

func checkIdentity0104(c *gin.Context) {
	client, err := getSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := CheckIdentity0104(client)
	c.JSON(http.StatusOK, result)
}

func checkIdentity0105(c *gin.Context) {
	client, err := getSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := CheckIdentity0105(client)
	c.JSON(http.StatusOK, result)
}

func checkIdentity0106(c *gin.Context) {
	client, err := getSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := CheckIdentity0106(client)
	c.JSON(http.StatusOK, result)
}

func checkIdentity0107(c *gin.Context) {
	client, err := getSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := CheckIdentity0107(client)
	c.JSON(http.StatusOK, result)
}

func checkIdentity0108(c *gin.Context) {
	client, err := getSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := CheckIdentity0108(client)
	c.JSON(http.StatusOK, result)
}
