package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gunh0/openstack-security-hub/checklist/keymanager"
	"github.com/gunh0/openstack-security-hub/util"
)

func RegisterKeyManagerRoutes(router *gin.RouterGroup) {
	router.GET("/check/key-manager-01-01", checkKeyManager0101)
}

// @Summary     Is the ownership of config files set to root/barbican? (/etc/barbican/barbican.conf)
// @Description Configuration files contain critical parameters and information required for smooth functioning of the component. If an unprivileged user, either intentionally or accidentally, modifies or deletes any of the parameters or the file itself then it would cause severe availability issues resulting in a denial of service to the other end users. User ownership of such critical configuration files must be set to root and group ownership must be set to barbican. Additionally, the containing directory should have the same ownership to ensure that new files are owned correctly.
// @Tags        Secrets Management
// @Produce     json
// @Success     200 {object} checklist.CheckResult
// @Router      /check/key-manager-01-01 [get]
func checkKeyManager0101(c *gin.Context) {
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

	result := keymanager.CheckKeyManager0101(client)
	c.JSON(http.StatusOK, result)
}
