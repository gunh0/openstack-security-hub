// checklist/dashboard/dashboard.go
package dashboard

import (
	"fmt"
	"strings"

	"github.com/gunh0/openstack-security-hub/checklist"
	"github.com/gunh0/openstack-security-hub/util"
	"golang.org/x/crypto/ssh"
)

// CheckDashboard01 checks if user/group ownership of config files is set to root/horizon
func CheckDashboard01(client *ssh.Client) checklist.CheckResult {
	const (
		description = "Is user/group of config files set to root/horizon?"
	)

	session, err := client.NewSession()
	if err != nil {
		return checklist.CheckResult{
			Description: description,
			Result:      "[ERROR]",
			Details:     fmt.Sprintf("Failed to create SSH session: %v", err),
		}
	}
	defer session.Close()

	// Check file permissions and ownership
	cmd := `
		if [ ! -f "/etc/openstack-dashboard/local_settings.py" ]; then
           echo "FILE_NOT_FOUND"
           exit 0
       	fi

		if [ ! -r "/etc/openstack-dashboard/local_settings.py" ]; then
           echo "PERMISSION_DENIED"
           exit 0
       	fi

		ownership=$(stat -L -c "%U %G" /etc/openstack-dashboard/local_settings.py 2>/dev/null)
	    echo "OWNERSHIP:$ownership"
	`

	output, err := session.CombinedOutput(cmd)
	result := strings.TrimSpace(string(output))

	// Process results
	switch {
	case strings.Contains(result, "PERMISSION_DENIED"):
		return checklist.CheckResult{
			Description: description,
			Result:      "[NA]",
			Details:     "Cannot check local_settings.py: permission denied",
		}
	case strings.Contains(result, "FILE_NOT_FOUND"):
		return checklist.CheckResult{
			Description: description,
			Result:      "[NA]",
			Details:     "local_settings.py not found",
		}
	}

	// Check ownership
	if strings.Contains(result, "OWNERSHIP:") {
		ownership := strings.TrimPrefix(result, "OWNERSHIP:")
		ownership = strings.TrimSpace(ownership)

		if ownership == "root horizon" {
			return checklist.CheckResult{
				Description: description,
				Result:      "[PASS]",
				Details:     "File ownership is correctly set to root:horizon",
			}
		}

		return checklist.CheckResult{
			Description: description,
			Result:      "[FAIL]",
			Details:     fmt.Sprintf("Current ownership is %s (expected: root horizon)", ownership),
		}
	}

	return checklist.CheckResult{
		Description: description,
		Result:      "[ERROR]",
		Details:     "Failed to determine file ownership",
	}
}

// CheckDashboard04 checks if CSRF_COOKIE_SECURE parameter is set to True
func CheckDashboard04(client *ssh.Client) checklist.CheckResult {
	return util.ExecuteScriptAndGetResult(
		client,
		"checklist/dashboard/dashboard-04.sh",
		"Is CSRF_COOKIE_SECURE parameter set to True?",
	)
}

// CheckDashboard05 checks if SESSION_COOKIE_SECURE parameter is set to True
func CheckDashboard05(client *ssh.Client) checklist.CheckResult {
	return util.ExecuteScriptAndGetResult(
		client,
		"checklist/dashboard/dashboard-05.sh",
		"Is SESSION_COOKIE_SECURE parameter set to True?",
	)
}
