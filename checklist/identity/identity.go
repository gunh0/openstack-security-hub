// checklist/identity/identity.go
package identity

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gunh0/openstack-security-hub/checklist"
	"golang.org/x/crypto/ssh"
)

// Common function to check file ownership
func checkFileOwnership(client *ssh.Client, filepath string) checklist.CheckResult {
	currentTime := time.Now().UTC().Format(time.RFC3339)
	session, err := client.NewSession()
	if err != nil {
		return checklist.CheckResult{
			Result:      "[ERROR]",
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
		return checklist.CheckResult{
			Result:      "[ERROR]",
			Description: fmt.Sprintf("Is user/group ownership of %s set to keystone?", filepath),
			Details:     fmt.Sprintf("Failed to execute command: %v", err),
			Timestamp:   currentTime,
		}
	}

	currentOwnership := strings.TrimSpace(string(output))
	if currentOwnership == "FILE_NOT_FOUND" {
		return checklist.CheckResult{
			Result:      "[NA]",
			Description: fmt.Sprintf("Is user/group ownership of %s set to keystone?", filepath),
			Details:     "File does not exist",
			Timestamp:   currentTime,
		}
	}

	// Check if ownership is correct
	if currentOwnership == "keystone keystone" {
		return checklist.CheckResult{
			Result:      "[PASS]",
			Description: fmt.Sprintf("Is user/group ownership of %s set to keystone?", filepath),
			Details:     fmt.Sprintf("Current ownership is correct: %s", currentOwnership),
			Timestamp:   currentTime,
		}
	}

	return checklist.CheckResult{
		Result:      "[FAIL]",
		Description: fmt.Sprintf("Is user/group ownership of %s set to keystone?", filepath),
		Details:     fmt.Sprintf("Current ownership: %s (expected: keystone keystone)", currentOwnership),
		Timestamp:   currentTime,
	}
}

// checkFilePermissions checks if file permissions are set correctly
func checkFilePermissions(client *ssh.Client, filepath string, isDirectory bool) checklist.CheckResult {
	currentTime := time.Now().UTC().Format(time.RFC3339)
	session, err := client.NewSession()
	if err != nil {
		return checklist.CheckResult{
			Result:      "[ERROR]",
			Description: fmt.Sprintf("Are strict permissions set for %s?", filepath),
			Details:     fmt.Sprintf("Failed to create SSH session: %v", err),
			Timestamp:   currentTime,
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
		return checklist.CheckResult{
			Result:      "[ERROR]",
			Description: fmt.Sprintf("Are strict permissions set for %s?", filepath),
			Details:     fmt.Sprintf("Failed to execute command: %v", err),
			Timestamp:   currentTime,
		}
	}

	currentPerms := strings.TrimSpace(string(output))
	if currentPerms == "FILE_NOT_FOUND" {
		return checklist.CheckResult{
			Result:      "[NA]",
			Description: fmt.Sprintf("Are strict permissions set for %s?", filepath),
			Details:     "File does not exist",
			Timestamp:   currentTime,
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
		return checklist.CheckResult{
			Result:      "[PASS]",
			Description: fmt.Sprintf("Are strict permissions set for %s?", filepath),
			Details:     fmt.Sprintf("Current permissions: %s (meets or exceeds required: %s)", currentPerms, expectedPerms),
			Timestamp:   currentTime,
		}
	}

	return checklist.CheckResult{
		Result:      "[FAIL]",
		Description: fmt.Sprintf("Are strict permissions set for %s?", filepath),
		Details:     fmt.Sprintf("Current permissions: %s (should be %s or stricter)", currentPerms, expectedPerms),
		Timestamp:   currentTime,
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
func CheckIdentity0101(client *ssh.Client) checklist.CheckResult {
	return checkFileOwnership(client, "/etc/keystone/keystone.conf")
}

func CheckIdentity0102(client *ssh.Client) checklist.CheckResult {
	return checkFileOwnership(client, "/etc/keystone/keystone-paste.ini")
}

func CheckIdentity0103(client *ssh.Client) checklist.CheckResult {
	return checkFileOwnership(client, "/etc/keystone/policy.json")
}

func CheckIdentity0104(client *ssh.Client) checklist.CheckResult {
	return checkFileOwnership(client, "/etc/keystone/logging.conf")
}

func CheckIdentity0105(client *ssh.Client) checklist.CheckResult {
	return checkFileOwnership(client, "/etc/keystone/ssl/certs/signing_cert.pem")
}

func CheckIdentity0106(client *ssh.Client) checklist.CheckResult {
	return checkFileOwnership(client, "/etc/keystone/ssl/private/signing_key.pem")
}

func CheckIdentity0107(client *ssh.Client) checklist.CheckResult {
	return checkFileOwnership(client, "/etc/keystone/ssl/certs/ca.pem")
}

func CheckIdentity0108(client *ssh.Client) checklist.CheckResult {
	return checkFileOwnership(client, "/etc/keystone")
}

// Identity-02 check functions
func CheckIdentity0201(client *ssh.Client) checklist.CheckResult {
	return checkFilePermissions(client, "/etc/keystone/keystone.conf", false)
}

func CheckIdentity0202(client *ssh.Client) checklist.CheckResult {
	return checkFilePermissions(client, "/etc/keystone/keystone-paste.ini", false)
}

func CheckIdentity0203(client *ssh.Client) checklist.CheckResult {
	return checkFilePermissions(client, "/etc/keystone/policy.json", false)
}

func CheckIdentity0204(client *ssh.Client) checklist.CheckResult {
	return checkFilePermissions(client, "/etc/keystone/logging.conf", false)
}

func CheckIdentity0205(client *ssh.Client) checklist.CheckResult {
	return checkFilePermissions(client, "/etc/keystone/ssl/certs/signing_cert.pem", false)
}

func CheckIdentity0206(client *ssh.Client) checklist.CheckResult {
	return checkFilePermissions(client, "/etc/keystone/ssl/private/signing_key.pem", false)
}

func CheckIdentity0207(client *ssh.Client) checklist.CheckResult {
	return checkFilePermissions(client, "/etc/keystone/ssl/certs/ca.pem", false)
}

func CheckIdentity0208(client *ssh.Client) checklist.CheckResult {
	return checkFilePermissions(client, "/etc/keystone", true)
}

func CheckIdentity03(client *ssh.Client) checklist.CheckResult {
	session, err := client.NewSession()
	if err != nil {
		return checklist.CheckResult{
			Result:      "[ERROR]",
			Description: "Is TLS enabled for Identity?",
			Details:     fmt.Sprintf("Failed to create SSH session: %v", err),
		}
	}
	defer session.Close()

	// Check if port 443 is in use (exact match)
	cmd := `netstat -tnlp 2>/dev/null | grep ':443 ' || echo "HTTPS_DISABLED"`

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return checklist.CheckResult{
			Result:      "[ERROR]",
			Description: "Is TLS enabled for Identity?",
			Details:     fmt.Sprintf("Failed to execute check: %v", err),
		}
	}

	result := strings.TrimSpace(string(output))
	if result == "HTTPS_DISABLED" {
		return checklist.CheckResult{
			Result:      "[FAIL]",
			Description: "Is TLS enabled for Identity?",
			Details:     "HTTPS port 443 is not in use",
		}
	}

	return checklist.CheckResult{
		Result:      "[PASS]",
		Description: "Is TLS enabled for Identity?",
		Details:     fmt.Sprintf("HTTPS port 443 is in use: %s", result),
	}
}

func CheckIdentity05(client *ssh.Client) checklist.CheckResult {
	const (
		description = "Is max_request_body_size set to default (114688)?"
		defaultSize = 114688
		maxSize     = 10485760 // 10MB
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

	// First check file access permission
	cmd := `
		if [ ! -r "/etc/keystone/keystone.conf" ]; then
			echo "PERMISSION_DENIED"
			exit 0
		fi
		
		if [ ! -f "/etc/keystone/keystone.conf" ]; then
			echo "FILE_NOT_FOUND"
			exit 0
		fi
 
		value=$(grep "^max_request_body_size" /etc/keystone/keystone.conf 2>/dev/null | awk -F "=" '{print $2}' | tr -d ' ')
		if [ -z "$value" ]; then
			echo "NOT_SET"
		else
			echo "SET:$value"
		fi
	`

	output, err := session.CombinedOutput(cmd)
	result := strings.TrimSpace(string(output))

	// Process results
	switch {
	case strings.Contains(result, "PERMISSION_DENIED"):
		return checklist.CheckResult{
			Description: description,
			Result:      "[NA]",
			Details:     "Cannot check keystone.conf: permission denied",
		}
	case strings.Contains(result, "FILE_NOT_FOUND"):
		return checklist.CheckResult{
			Description: description,
			Result:      "[NA]",
			Details:     "Keystone configuration file not found",
		}
	case strings.Contains(result, "NOT_SET"):
		return checklist.CheckResult{
			Description: description,
			Result:      "[FAIL]",
			Details:     "max_request_body_size parameter is not set in keystone.conf",
		}
	}

	// Process SET value
	value := strings.TrimPrefix(result, "SET:")
	if value == strconv.Itoa(defaultSize) {
		return checklist.CheckResult{
			Description: description,
			Result:      "[PASS]",
			Details:     "max_request_body_size is set to the default value (114688)",
		}
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return checklist.CheckResult{
			Description: description,
			Result:      "[NA]",
			Details:     fmt.Sprintf("Unable to parse max_request_body_size value: %s", value),
		}
	}

	// Check if value is within reasonable range
	if intValue >= defaultSize && intValue <= maxSize {
		return checklist.CheckResult{
			Description: description,
			Result:      "[PASS]",
			Details:     fmt.Sprintf("max_request_body_size is set to a reasonable value: %s bytes", value),
		}
	}

	return checklist.CheckResult{
		Description: description,
		Result:      "[FAIL]",
		Details:     fmt.Sprintf("max_request_body_size is set to a potentially unsafe value: %s bytes", value),
	}
}

func CheckIdentity06(client *ssh.Client) checklist.CheckResult {
	const (
		description = "Disable admin token in /etc/keystone/keystone.conf"
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

	// First check keystone.conf and extract admin_token value
	cmd := `
		# Check keystone.conf first
		if [ ! -r "/etc/keystone/keystone.conf" ]; then
			echo "KEYSTONE_CONF_PERMISSION_DENIED"
			exit 0
		fi
		
		if [ ! -f "/etc/keystone/keystone.conf" ]; then
			echo "KEYSTONE_CONF_NOT_FOUND"
			exit 0
		fi
 
		# Check admin_token in keystone.conf
		admin_token=$(grep "^admin_token" /etc/keystone/keystone.conf 2>/dev/null | awk -F "=" '{print $2}' | tr -d ' ')
		echo "ADMIN_TOKEN:${admin_token:-NOTSET}"
 
		# Check keystone-paste.ini only if it exists
		if [ -f "/etc/keystone/keystone-paste.ini" ]; then
			if [ ! -r "/etc/keystone/keystone-paste.ini" ]; then
				echo "PASTE_INI_PERMISSION_DENIED"
			else
				auth_middleware=$(grep "AdminTokenAuthMiddleware" /etc/keystone/keystone-paste.ini 2>/dev/null)
				if [ -n "$auth_middleware" ]; then
					echo "AUTH_MIDDLEWARE:EXISTS"
				else
					echo "AUTH_MIDDLEWARE:NOTFOUND"
				fi
			fi
		else
			echo "PASTE_INI_NOT_FOUND"
		fi
	`

	output, err := session.CombinedOutput(cmd)
	result := strings.TrimSpace(string(output))
	lines := strings.Split(result, "\n")

	// Process results
	if strings.Contains(result, "KEYSTONE_CONF_PERMISSION_DENIED") {
		return checklist.CheckResult{
			Description: description,
			Result:      "[NA]",
			Details:     "Cannot check keystone.conf: permission denied",
		}
	}

	if strings.Contains(result, "KEYSTONE_CONF_NOT_FOUND") {
		return checklist.CheckResult{
			Description: description,
			Result:      "[NA]",
			Details:     "keystone.conf not found",
		}
	}

	var details strings.Builder
	var adminTokenDisabled bool
	var middlewareDisabled bool = true // Default to true if paste.ini doesn't exist

	// Process admin_token status
	for _, line := range lines {
		if strings.HasPrefix(line, "ADMIN_TOKEN:") {
			value := strings.TrimPrefix(line, "ADMIN_TOKEN:")
			adminTokenDisabled = value == "NOTSET" || value == "<none>"
			if !adminTokenDisabled {
				details.WriteString(fmt.Sprintf("- admin_token is set with value: %s\n", value))
			}
		}

		// Process middleware status if paste.ini exists
		if strings.HasPrefix(line, "AUTH_MIDDLEWARE:") {
			switch strings.TrimPrefix(line, "AUTH_MIDDLEWARE:") {
			case "EXISTS":
				middlewareDisabled = false
				details.WriteString("- AdminTokenAuthMiddleware is present in keystone-paste.ini\n")
			case "NOTFOUND":
				middlewareDisabled = true
			}
		}
	}

	// If paste.ini doesn't exist or has permission issues, add note to details
	if strings.Contains(result, "PASTE_INI_NOT_FOUND") {
		details.WriteString("Note: keystone-paste.ini not found (this is acceptable)\n")
	} else if strings.Contains(result, "PASTE_INI_PERMISSION_DENIED") {
		details.WriteString("Warning: Cannot read keystone-paste.ini due to permission denied\n")
	}

	// Determine final result
	switch {
	case adminTokenDisabled && middlewareDisabled:
		if details.Len() == 0 {
			details.WriteString("admin_token is disabled and AdminTokenAuthMiddleware is not present")
		}
		return checklist.CheckResult{
			Description: description,
			Result:      "[PASS]",
			Details:     details.String(),
		}
	default:
		if details.Len() == 0 {
			details.WriteString("Security configuration issues found")
		}
		return checklist.CheckResult{
			Description: description,
			Result:      "[FAIL]",
			Details:     details.String(),
		}
	}
}
