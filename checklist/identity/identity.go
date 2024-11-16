// checklist/identity/identity.go
package identity

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

// CheckResult represents a check result
type CheckResult struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Details     string `json:"details"`
}

// Common function to check file ownership
func checkFileOwnership(client *ssh.Client, filepath string) CheckResult {
	session, err := client.NewSession()
	if err != nil {
		return CheckResult{
			Status:      "[ERROR]",
			Description: fmt.Sprintf("Is user/group ownership of %s set to keystone?", filepath),
			Details:     fmt.Sprintf("Failed to create SSH session: %v", err),
		}
	}
	defer session.Close()

	// First check current ownership
	cmd := fmt.Sprintf(`
        if [ -e "%s" ]; then
            stat -L -c "%%U %%G" "%s"
        else
            echo "FILE_NOT_FOUND"
        fi
    `, filepath, filepath)

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return CheckResult{
			Status:      "[ERROR]",
			Description: fmt.Sprintf("Is user/group ownership of %s set to keystone?", filepath),
			Details:     fmt.Sprintf("Failed to execute command: %v", err),
		}
	}

	currentOwnership := strings.TrimSpace(string(output))
	if currentOwnership == "FILE_NOT_FOUND" {
		return CheckResult{
			Status:      "[NA]",
			Description: fmt.Sprintf("Is user/group ownership of %s set to keystone?", filepath),
			Details:     "File does not exist",
		}
	}

	// Check if ownership is correct
	if currentOwnership == "keystone keystone" {
		return CheckResult{
			Status:      "[PASS]",
			Description: fmt.Sprintf("Is user/group ownership of %s set to keystone?", filepath),
			Details:     fmt.Sprintf("Current ownership is correct: %s", currentOwnership),
		}
	}

	return CheckResult{
		Status:      "[FAIL]",
		Description: fmt.Sprintf("Is user/group ownership of %s set to keystone?", filepath),
		Details:     fmt.Sprintf("Current ownership: %s (expected: keystone keystone)", currentOwnership),
	}
}

// Individual check functions
func CheckIdentity0101(client *ssh.Client) CheckResult {
	return checkFileOwnership(client, "/etc/keystone/keystone.conf")
}

func CheckIdentity0102(client *ssh.Client) CheckResult {
	return checkFileOwnership(client, "/etc/keystone/keystone-paste.ini")
}

func CheckIdentity0103(client *ssh.Client) CheckResult {
	return checkFileOwnership(client, "/etc/keystone/policy.json")
}

func CheckIdentity0104(client *ssh.Client) CheckResult {
	return checkFileOwnership(client, "/etc/keystone/logging.conf")
}

func CheckIdentity0105(client *ssh.Client) CheckResult {
	return checkFileOwnership(client, "/etc/keystone/ssl/certs/signing_cert.pem")
}

func CheckIdentity0106(client *ssh.Client) CheckResult {
	return checkFileOwnership(client, "/etc/keystone/ssl/private/signing_key.pem")
}

func CheckIdentity0107(client *ssh.Client) CheckResult {
	return checkFileOwnership(client, "/etc/keystone/ssl/certs/ca.pem")
}

func CheckIdentity0108(client *ssh.Client) CheckResult {
	return checkFileOwnership(client, "/etc/keystone")
}
