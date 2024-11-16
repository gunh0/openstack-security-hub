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

	// Execute the stat command with grep for keystone ownership
	cmd := fmt.Sprintf(`stat -L -c "%%U %%G" %s | egrep "keystone keystone" || echo "FAILED"`, filepath)
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		// Check if the error is due to file not existing
		checkExist := fmt.Sprintf("stat -L -c \"%%U %%G\" %s", filepath)
		_, existErr := session.CombinedOutput(checkExist)
		if existErr != nil {
			return CheckResult{
				Status:      "[NA]",
				Description: fmt.Sprintf("Is user/group ownership of %s set to keystone?", filepath),
				Details:     "File does not exist",
			}
		}
	}

	result := strings.TrimSpace(string(output))
	if result == "FAILED" {
		// File exists but ownership is not keystone:keystone
		// Get current ownership for detailed output
		cmd = fmt.Sprintf(`stat -L -c "%%U %%G" %s`, filepath)
		currentOwnership, _ := session.CombinedOutput(cmd)
		return CheckResult{
			Status:      "[FAIL]",
			Description: fmt.Sprintf("Is user/group ownership of %s set to keystone?", filepath),
			Details:     fmt.Sprintf("Current ownership: %s (expected: keystone keystone)", strings.TrimSpace(string(currentOwnership))),
		}
	}

	if result == "keystone keystone" {
		return CheckResult{
			Status:      "[PASS]",
			Description: fmt.Sprintf("Is user/group ownership of %s set to keystone?", filepath),
			Details:     "File is owned by keystone:keystone",
		}
	}

	return CheckResult{
		Status:      "[ERROR]",
		Description: fmt.Sprintf("Is user/group ownership of %s set to keystone?", filepath),
		Details:     "Unexpected error while checking file ownership",
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
