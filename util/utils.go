// util/utils.go
package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gunh0/openstack-security-hub/checklist"
	"golang.org/x/crypto/ssh"
)

// GetSSHClient returns a new SSH client using environment variables for configuration
func GetSSHClient() (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: os.Getenv("SSH_USER"),
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("SSH_PASSWORD")),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return ssh.Dial("tcp", os.Getenv("SSH_HOST"), config)
}

// PrettyPrintResult prints a formatted check result with clear visual separation
func PrettyPrintResult(result checklist.CheckResult) {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("Description: %s\n", result.Description)
	fmt.Printf("Result: %s\n", result.Result)
	fmt.Printf("Details: %s\n", result.Details)
	fmt.Printf("Timestamp: %s\n", result.Timestamp)
	fmt.Println(strings.Repeat("-", 100))
}

// SSHClient wraps an ssh.Client to provide additional functionality
type SSHClient struct {
	client *ssh.Client
}

// ExecuteScript executes a shell script on the remote host and returns the output
func (c *SSHClient) ExecuteScript(scriptPath string) (string, error) {
	// Read the local script file into memory
	content, err := os.ReadFile(scriptPath)
	if err != nil {
		return "", fmt.Errorf("failed to read script: %v", err)
	}

	// Create a new SSH session for script execution
	session, err := c.client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	// Set up output and error buffers
	var outputBuffer, errorBuffer bytes.Buffer
	session.Stdout = &outputBuffer
	session.Stderr = &errorBuffer

	// Execute the script using heredoc to handle multiline scripts
	err = session.Run(fmt.Sprintf("bash -s << 'EOF'\n%s\nEOF", string(content)))
	if err != nil {
		if errorBuffer.Len() > 0 {
			return "", fmt.Errorf("script execution failed: %v, stderr: %s", err, errorBuffer.String())
		}
		return "", fmt.Errorf("script execution failed: %v", err)
	}

	return outputBuffer.String(), nil
}

// Close safely terminates the SSH connection
func (c *SSHClient) Close() error {
	return c.client.Close()
}

// ExecuteScriptAndGetResult executes a shell script via SSH and returns the parsed CheckResult.
// It captures and displays the full script output, then extracts and parses the JSON result
// from the last line of output.
func ExecuteScriptAndGetResult(client *ssh.Client, scriptPath string, description string) checklist.CheckResult {
	// Validate SSH client initialization
	if client == nil {
		return checklist.CheckResult{
			Description: description,
			Result:      "[ERROR]",
			Details:     "SSH client is nil",
		}
	}
	fmt.Printf("[SSH] Connected to: %v\n", client.RemoteAddr())

	// Create new SSH session
	session, err := client.NewSession()
	if err != nil {
		fmt.Printf("[ERROR] Failed to create SSH session: %v\n", err)
		return checklist.CheckResult{
			Description: description,
			Result:      "[ERROR]",
			Details:     fmt.Sprintf("Failed to create SSH session: %v", err),
		}
	}
	defer session.Close()

	// Get and validate script path
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("[ERROR] Failed to get working directory: %v\n", err)
		return checklist.CheckResult{
			Description: description,
			Result:      "[ERROR]",
			Details:     fmt.Sprintf("Failed to get working directory: %v", err),
		}
	}

	fullScriptPath := filepath.Join(pwd, scriptPath)
	fmt.Printf("[INFO] Executing script: %s\n", fullScriptPath)

	// Check script existence
	if _, err := os.Stat(fullScriptPath); os.IsNotExist(err) {
		fmt.Printf("[ERROR] Script not found: %s\n", fullScriptPath)
		return checklist.CheckResult{
			Description: description,
			Result:      "[ERROR]",
			Details:     fmt.Sprintf("Script not found at path: %s", fullScriptPath),
		}
	}

	// Read script content
	scriptContent, err := os.ReadFile(fullScriptPath)
	if err != nil {
		fmt.Printf("[ERROR] Failed to read script: %v\n", err)
		return checklist.CheckResult{
			Description: description,
			Result:      "[ERROR]",
			Details:     fmt.Sprintf("Failed to read script: %v", err),
		}
	}

	// Execute script and capture output directly
	output, err := session.CombinedOutput(string(scriptContent))
	if err != nil {
		fmt.Printf("[ERROR] Script execution failed: %v\n", err)
		return checklist.CheckResult{
			Description: description,
			Result:      "[ERROR]",
			Details:     fmt.Sprintf("Script execution failed: %v\nOutput: %s", err, string(output)),
		}
	}

	// Display full script output
	fmt.Printf("[SCRIPT OUTPUT]\n%s", string(output))

	// Extract JSON from last line and parse result
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 0 {
		return checklist.CheckResult{
			Description: description,
			Result:      "[ERROR]",
			Details:     "No output from script",
		}
	}

	jsonLine := lines[len(lines)-1]
	var result checklist.CheckResult
	if err := json.Unmarshal([]byte(jsonLine), &result); err != nil {
		return checklist.CheckResult{
			Description: description,
			Result:      "[ERROR]",
			Details:     fmt.Sprintf("Failed to parse JSON result: %v", err),
		}
	}

	return result
}
