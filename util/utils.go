package util

import (
	"fmt"
	"os"
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
