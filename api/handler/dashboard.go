package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gunh0/openstack-security-hub/checklist/dashboard"
	"github.com/gunh0/openstack-security-hub/util"
)

func RegisterDashboardRoutes(router *gin.RouterGroup) {
	router.GET("/check/dashboard-01", handleDashboard01)
}

// @Summary	 Is user/group ownership of config files set to root/horizon?Is user/group of config files set to root/horizon?
// @Description Configuration files contain critical parameters and information required for smooth functioning of the component. If an unprivileged user, either intentionally or accidentally modifies or deletes any of the parameters or the file itself then it would cause severe availability issues causing a denial of service to the other end users. Thus user ownership of such critical configuration files must be set to root and group ownership must be set to horizon.
// @Tags		dashboard
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
