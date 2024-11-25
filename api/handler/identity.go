package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gunh0/openstack-security-hub/checklist"
	"github.com/gunh0/openstack-security-hub/checklist/identity"
	"github.com/gunh0/openstack-security-hub/util"
	"golang.org/x/crypto/ssh"
)

// RegisterIdentityRoutes registers all identity check routes
func RegisterIdentityRoutes(router *gin.RouterGroup) {
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

// @Summary     Run all Identity-01 security checks
// @Description Execute all Identity-01 series security checks for OpenStack
// @Tags        identity
// @Produce     json
// @Success     200 {array}  checklist.CheckResult
// @Failure     500 {object} map[string]string
// @Router      /check/identity-01 [get]
func handleIdentity01(c *gin.Context) {
	client, err := util.GetSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer client.Close()

	checks := []struct {
		name string
		fn   func(*ssh.Client) checklist.CheckResult
	}{
		{"Identity-01-01", identity.CheckIdentity0101},
		{"Identity-01-02", identity.CheckIdentity0102},
		{"Identity-01-03", identity.CheckIdentity0103},
		{"Identity-01-04", identity.CheckIdentity0104},
		{"Identity-01-05", identity.CheckIdentity0105},
		{"Identity-01-06", identity.CheckIdentity0106},
		{"Identity-01-07", identity.CheckIdentity0107},
		{"Identity-01-08", identity.CheckIdentity0108},
	}

	var results []map[string]checklist.CheckResult
	for _, check := range checks {
		result := check.fn(client)
		results = append(results, map[string]checklist.CheckResult{check.name: result})
	}

	c.JSON(http.StatusOK, results)
}

// Utility function to reduce duplication
func handleIdentityCheck(c *gin.Context, checkFn func(*ssh.Client) checklist.CheckResult) {
	client, err := util.GetSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := checkFn(client)
	c.JSON(http.StatusOK, result)
}

// @Summary     Run Identity-01-01 check
// @Description Check Identity configuration security
// @Tags        identity
// @Produce     json
// @Success     200 {object} checklist.CheckResult
// @Failure     500 {object} map[string]string
// @Router      /check/identity-01-01 [get]
func checkIdentity0101(c *gin.Context) {
	client, err := util.GetSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0101(client)
	c.JSON(http.StatusOK, result)
}

// @Summary     Run Identity-01-02 check
// @Description Check Identity configuration security
// @Tags        identity
// @Produce     json
// @Success     200 {object} checklist.CheckResult
// @Failure     500 {object} map[string]string
// @Router      /check/identity-01-02 [get]
func checkIdentity0102(c *gin.Context) {
	client, err := util.GetSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0102(client)
	c.JSON(http.StatusOK, result)
}

func checkIdentity0103(c *gin.Context) {
	client, err := util.GetSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0103(client)
	c.JSON(http.StatusOK, result)
}

func checkIdentity0104(c *gin.Context) {
	client, err := util.GetSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0104(client)
	c.JSON(http.StatusOK, result)
}

func checkIdentity0105(c *gin.Context) {
	client, err := util.GetSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0105(client)
	c.JSON(http.StatusOK, result)
}

func checkIdentity0106(c *gin.Context) {
	client, err := util.GetSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0106(client)
	c.JSON(http.StatusOK, result)
}

func checkIdentity0107(c *gin.Context) {
	client, err := util.GetSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0107(client)
	c.JSON(http.StatusOK, result)
}

func checkIdentity0108(c *gin.Context) {
	client, err := util.GetSSHClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0108(client)
	c.JSON(http.StatusOK, result)
}
