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

// PrettyPrintResult prints a formatted check result
func PrettyPrintResult(result checklist.CheckResult) {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("Description: %s\n", result.Description)
	fmt.Printf("Result: %s\n", result.Result)
	fmt.Printf("Details: %s\n", result.Details)
	fmt.Println(strings.Repeat("-", 100))
}

type SSHClient struct {
	client *ssh.Client
}

// NewSSHClient creates a new SSH client using environment variables
func NewSSHClient() (*SSHClient, error) {
	config := &ssh.ClientConfig{
		User: os.Getenv("SSH_USER"),
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("SSH_PASSWORD")),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", os.Getenv("SSH_HOST"), config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %v", err)
	}

	return &SSHClient{client: client}, nil
}

// ExecuteScript executes a script file on the remote host
func (c *SSHClient) ExecuteScript(scriptPath string) (string, error) {
	// Read local script file
	content, err := os.ReadFile(scriptPath)
	if err != nil {
		return "", fmt.Errorf("failed to read script: %v", err)
	}

	// Create new session
	session, err := c.client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	// Create buffers for output
	var outputBuffer, errorBuffer bytes.Buffer
	session.Stdout = &outputBuffer
	session.Stderr = &errorBuffer

	// Execute script directly using bash
	err = session.Run(fmt.Sprintf("bash -s << 'EOF'\n%s\nEOF", string(content)))
	if err != nil {
		if errorBuffer.Len() > 0 {
			return "", fmt.Errorf("script execution failed: %v, stderr: %s", err, errorBuffer.String())
		}
		return "", fmt.Errorf("script execution failed: %v", err)
	}

	return outputBuffer.String(), nil
}

// Close closes the SSH connection
func (c *SSHClient) Close() error {
	return c.client.Close()
}

// ExecuteScriptAndGetResult executes a shell script and returns the parsed CheckResult
func ExecuteScriptAndGetResult(client *ssh.Client, scriptPath string, description string) checklist.CheckResult {
	session, err := client.NewSession()
	if err != nil {
		return checklist.CheckResult{
			Description: description,
			Result:      "[ERROR]",
			Details:     fmt.Sprintf("Failed to create SSH session: %v", err),
		}
	}
	defer session.Close()

	// Get absolute path and check if script exists
	pwd, err := os.Getwd()
	if err != nil {
		return checklist.CheckResult{
			Description: description,
			Result:      "[ERROR]",
			Details:     fmt.Sprintf("Failed to get working directory: %v", err),
		}
	}

	fullScriptPath := filepath.Join(pwd, scriptPath)
	if _, err := os.Stat(fullScriptPath); os.IsNotExist(err) {
		return checklist.CheckResult{
			Description: description,
			Result:      "[ERROR]",
			Details:     fmt.Sprintf("Script not found at path: %s", fullScriptPath),
		}
	}

	// Read script content
	scriptContent, err := os.ReadFile(fullScriptPath)
	if err != nil {
		return checklist.CheckResult{
			Description: description,
			Result:      "[ERROR]",
			Details:     fmt.Sprintf("Failed to read script: %v", err),
		}
	}

	// Execute script and capture output
	var outputBuf bytes.Buffer
	session.Stdout = &outputBuf
	session.Stderr = &outputBuf

	err = session.Run(string(scriptContent)) // Changed from := to =
	outputStr := outputBuf.String()

	if err != nil {
		return checklist.CheckResult{
			Description: description,
			Result:      "[ERROR]",
			Details:     fmt.Sprintf("Script execution failed: %v\nOutput: %s", err, outputStr),
		}
	}

	// Debug output (optional)
	fmt.Printf("Script path: %s\n", fullScriptPath)
	fmt.Printf("Script output:\n%s", outputStr)

	// Parse JSON result
	startIdx := strings.LastIndex(outputStr, "{")
	endIdx := strings.LastIndex(outputStr, "}")

	if startIdx == -1 || endIdx == -1 || startIdx > endIdx {
		return checklist.CheckResult{
			Description: description,
			Result:      "[ERROR]",
			Details:     fmt.Sprintf("Failed to locate JSON in script output: %s", outputStr),
		}
	}

	jsonStr := outputStr[startIdx : endIdx+1]
	var result checklist.CheckResult
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return checklist.CheckResult{
			Description: description,
			Result:      "[ERROR]",
			Details:     fmt.Sprintf("Failed to parse JSON result: %v\nOutput: %s", err, outputStr),
		}
	}

	return result
}
