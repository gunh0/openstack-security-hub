package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gunh0/openstack-security-hub/checklist/dashboard"
	"github.com/gunh0/openstack-security-hub/util"
)

func RegisterDashboardRoutes(router *gin.RouterGroup) {
	router.GET("/check/dashboard-01", handleDashboard01)
	router.GET("/check/dashboard-04", handleDashboard04)
	router.GET("/check/dashboard-05", handleDashboard05)
	router.GET("/check/dashboard-06", handleDashboard06)
}

// @Summary	 Is user/group ownership of config files set to root/horizon?Is user/group of config files set to root/horizon?
// @Description Configuration files contain critical parameters and information required for smooth functioning of the component. If an unprivileged user, either intentionally or accidentally modifies or deletes any of the parameters or the file itself then it would cause severe availability issues causing a denial of service to the other end users. Thus user ownership of such critical configuration files must be set to root and group ownership must be set to horizon.
// @Tags		Dashboard
// @Produce		json
// @Success		200	{array}	checklist.CheckResult
// @Router		/check/dashboard-01	[get]
func handleDashboard01(c *gin.Context) {
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

	result := dashboard.CheckDashboard01(client)
	c.JSON(http.StatusOK, result)
}

// @Summary Is CSRF_COOKIE_SECURE parameter set to True?
// @Description CSRF (Cross-site request forgery) is an attack which forces an end user to execute unauthorized commands on a web application in which he/she is currently authenticated. A successful CSRF exploit can compromise end user data and operations. If the targeted end user has admin privileges, this can compromise the entire web application.
// @Tags Dashboard
// @Produce json
// @Success 200 {array} checklist.CheckResult
// @Router /check/dashboard-04 [get]
func handleDashboard04(c *gin.Context) {
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

	result := dashboard.CheckDashboard04(client)
	c.JSON(http.StatusOK, result)
}

// @Summary Is SESSION_COOKIE_SECURE parameter set to True?
// @Description The “SECURE” cookie attribute instructs web browsers to only send the cookie through an encrypted HTTPS (SSL/TLS) connection. This session protection mechanism is mandatory to prevent the disclosure of the session ID through MitM (Man-in-the-Middle) attacks. It ensures that an attacker cannot simply capture the session ID from web browser traffic.
// @Tags Dashboard
// @Produce json
// @Success 200 {array} checklist.CheckResult
// @Router /check/dashboard-05 [get]
func handleDashboard05(c *gin.Context) {
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

	result := dashboard.CheckDashboard05(client)
	c.JSON(http.StatusOK, result)
}

// @Summary Is SESSION_COOKIE_HTTPONLY parameter set to True?
// @Description The “HTTPONLY” cookie attribute instructs web browsers not to allow scripts (e.g. JavaScript or VBscript) an ability to access the cookies via the DOM document.cookie object. This session ID protection is mandatory to prevent session ID stealing through XSS attacks.
// @Tags Dashboard
// @Produce json
// @Success 200 {array} checklist.CheckResult
// @Router /check/dashboard-06 [get]
func handleDashboard06(c *gin.Context) {
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

	result := dashboard.CheckDashboard06(client)
	c.JSON(http.StatusOK, result)
}
