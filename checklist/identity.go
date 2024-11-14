// checklist/identity.go
package checklist

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

type CheckResult struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Details     string `json:"details"`
}

func CheckIdentity01(client *ssh.Client) CheckResult {
	// Simple installation check - just verify directory and config file existence
	checkInstallCmd := `
        if [ -d "/etc/keystone" ] && [ -f "/etc/keystone/keystone.conf" ]; then
            echo "INSTALLED"
        else
            echo "NOT_INSTALLED"
        fi
    `

	session, err := client.NewSession()
	if err != nil {
		return CheckResult{
			Status:      "[ERROR]",
			Description: "Failed to create SSH session",
			Details:     err.Error(),
		}
	}

	installStatus, err := session.CombinedOutput(checkInstallCmd)
	session.Close()

	if err != nil {
		return CheckResult{
			Status:      "[ERROR]",
			Description: "Failed to check Keystone installation status",
			Details:     err.Error(),
		}
	}

	statusStr := strings.TrimSpace(string(installStatus))
	if statusStr == "NOT_INSTALLED" {
		return CheckResult{
			Status:      "[NA]",
			Description: "Keystone is not installed",
			Details:     "Please install Keystone before running this check",
		}
	}

	// File ownership check
	cmd := `
        files_to_check=(
            "/etc/keystone/keystone.conf"
            "/etc/keystone/keystone-uwsgi-public.ini"
            "/etc/keystone/credential-keys"
            "/etc/keystone/fernet-keys"
            "/etc/keystone"
        )

        for file in "${files_to_check[@]}"; do
            if [ -e "$file" ]; then
                stat -L -c "%U %G %n" "$file"
            fi
        done
    `

	session, err = client.NewSession()
	if err != nil {
		return CheckResult{
			Status:      "[ERROR]",
			Description: "Failed to create SSH session",
			Details:     err.Error(),
		}
	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return CheckResult{
			Status:      "[ERROR]",
			Description: "Failed to execute file check command",
			Details:     err.Error(),
		}
	}

	lines := strings.Split(string(output), "\n")
	var incorrectFiles []string

	for _, line := range lines {
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) >= 2 {
			user := fields[0]
			group := fields[1]
			path := strings.Join(fields[2:], " ")

			if user != "keystone" || group != "keystone" {
				incorrectFiles = append(incorrectFiles, fmt.Sprintf("%s (current: %s %s)", path, user, group))
			}
		}
	}

	if len(incorrectFiles) == 0 {
		return CheckResult{
			Status:      "[PASS]",
			Description: "All configuration files are properly owned by keystone user and group",
			Details:     string(output),
		}
	}

	return CheckResult{
		Status:      "[FAIL]",
		Description: "Some configuration files have incorrect ownership",
		Details:     "Files with incorrect ownership:\n" + strings.Join(incorrectFiles, "\n"),
	}
}
