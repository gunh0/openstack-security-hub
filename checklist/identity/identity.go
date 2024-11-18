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

// checkFilePermissions checks if file permissions are set correctly
func checkFilePermissions(client *ssh.Client, filepath string, isDirectory bool) CheckResult {
	session, err := client.NewSession()
	if err != nil {
		return CheckResult{
			Status:      "[ERROR]",
			Description: fmt.Sprintf("Are strict permissions set for %s?", filepath),
			Details:     fmt.Sprintf("Failed to create SSH session: %v", err),
		}
	}
	defer session.Close()

	// Check current permissions
	cmd := fmt.Sprintf(`
        if [ -e "%s" ]; then
            stat -L -c "%%a" "%s"
        else
            echo "FILE_NOT_FOUND"
        fi
    `, filepath, filepath)

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return CheckResult{
			Status:      "[ERROR]",
			Description: fmt.Sprintf("Are strict permissions set for %s?", filepath),
			Details:     fmt.Sprintf("Failed to execute command: %v", err),
		}
	}

	currentPerms := strings.TrimSpace(string(output))
	if currentPerms == "FILE_NOT_FOUND" {
		return CheckResult{
			Status:      "[NA]",
			Description: fmt.Sprintf("Are strict permissions set for %s?", filepath),
			Details:     "File does not exist",
		}
	}

	// Check if permissions are correct
	// For directories, check if permissions are 750 or stricter
	// For files, check if permissions are 640 or stricter
	var isValid bool
	var expectedPerms string

	if isDirectory {
		expectedPerms = "750"
		permsInt := parseOctal(currentPerms)
		isValid = permsInt <= parseOctal("750")
	} else {
		expectedPerms = "640"
		permsInt := parseOctal(currentPerms)
		isValid = permsInt <= parseOctal("640")
	}

	if isValid {
		return CheckResult{
			Status:      "[PASS]",
			Description: fmt.Sprintf("Are strict permissions set for %s?", filepath),
			Details:     fmt.Sprintf("Current permissions: %s (meets or exceeds required: %s)", currentPerms, expectedPerms),
		}
	}

	return CheckResult{
		Status:      "[FAIL]",
		Description: fmt.Sprintf("Are strict permissions set for %s?", filepath),
		Details:     fmt.Sprintf("Current permissions: %s (should be %s or stricter)", currentPerms, expectedPerms),
	}
}

// Helper function to parse octal permissions
func parseOctal(s string) int {
	n := 0
	for _, c := range s {
		n = n*8 + int(c-'0')
	}
	return n
}

// Identity-01 check functions
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

// Identity-02 check functions
func CheckIdentity0201(client *ssh.Client) CheckResult {
	return checkFilePermissions(client, "/etc/keystone/keystone.conf", false)
}

func CheckIdentity0202(client *ssh.Client) CheckResult {
	return checkFilePermissions(client, "/etc/keystone/keystone-paste.ini", false)
}

func CheckIdentity0203(client *ssh.Client) CheckResult {
	return checkFilePermissions(client, "/etc/keystone/policy.json", false)
}

func CheckIdentity0204(client *ssh.Client) CheckResult {
	return checkFilePermissions(client, "/etc/keystone/logging.conf", false)
}

func CheckIdentity0205(client *ssh.Client) CheckResult {
	return checkFilePermissions(client, "/etc/keystone/ssl/certs/signing_cert.pem", false)
}

func CheckIdentity0206(client *ssh.Client) CheckResult {
	return checkFilePermissions(client, "/etc/keystone/ssl/private/signing_key.pem", false)
}

func CheckIdentity0207(client *ssh.Client) CheckResult {
	return checkFilePermissions(client, "/etc/keystone/ssl/certs/ca.pem", false)
}

func CheckIdentity0208(client *ssh.Client) CheckResult {
	return checkFilePermissions(client, "/etc/keystone", true)
}

func CheckIdentity03(client *ssh.Client) CheckResult {
	session, err := client.NewSession()
	if err != nil {
		return CheckResult{
			Status:      "[ERROR]",
			Description: "Is TLS enabled for Identity?",
			Details:     fmt.Sprintf("Failed to create SSH session: %v", err),
		}
	}
	defer session.Close()

	// Check if port 443 is in use (exact match)
	cmd := `netstat -tnlp 2>/dev/null | grep ':443 ' || echo "HTTPS_DISABLED"`

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return CheckResult{
			Status:      "[ERROR]",
			Description: "Is TLS enabled for Identity?",
			Details:     fmt.Sprintf("Failed to execute check: %v", err),
		}
	}

	result := strings.TrimSpace(string(output))
	if result == "HTTPS_DISABLED" {
		return CheckResult{
			Status:      "[FAIL]",
			Description: "Is TLS enabled for Identity?",
			Details:     "HTTPS port 443 is not in use",
		}
	}

	return CheckResult{
		Status:      "[PASS]",
		Description: "Is TLS enabled for Identity?",
		Details:     fmt.Sprintf("HTTPS port 443 is in use: %s", result),
	}
}
