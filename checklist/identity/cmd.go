// checklist/identity/cmd.go
package identity

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

// InitCommands initializes all identity-related commands
func InitCommands(rootCmd *cobra.Command) {
	identity01Cmd := &cobra.Command{
		Use:   "identity-01",
		Short: "Run all identity-01 checks",
		Run:   runAllIdentity01Checks,
	}

	identity0101Cmd := &cobra.Command{
		Use:   "identity-01-01",
		Short: "Check keystone.conf ownership",
		Run:   runIdentity0101Check,
	}

	identity0102Cmd := &cobra.Command{
		Use:   "identity-01-02",
		Short: "Check keystone-paste.ini ownership",
		Run:   runIdentity0102Check,
	}

	identity0103Cmd := &cobra.Command{
		Use:   "identity-01-03",
		Short: "Check policy.json ownership",
		Run:   runIdentity0103Check,
	}

	identity0104Cmd := &cobra.Command{
		Use:   "identity-01-04",
		Short: "Check logging.conf ownership",
		Run:   runIdentity0104Check,
	}

	identity0105Cmd := &cobra.Command{
		Use:   "identity-01-05",
		Short: "Check signing_cert.pem ownership",
		Run:   runIdentity0105Check,
	}

	identity0106Cmd := &cobra.Command{
		Use:   "identity-01-06",
		Short: "Check signing_key.pem ownership",
		Run:   runIdentity0106Check,
	}

	identity0107Cmd := &cobra.Command{
		Use:   "identity-01-07",
		Short: "Check ca.pem ownership",
		Run:   runIdentity0107Check,
	}

	identity0108Cmd := &cobra.Command{
		Use:   "identity-01-08",
		Short: "Check /etc/keystone directory ownership",
		Run:   runIdentity0108Check,
	}

	// Add Identity-02 command
	identity02Cmd := &cobra.Command{
		Use:   "identity-02",
		Short: "Run all identity-02 checks",
		Run:   runAllIdentity02Checks,
	}

	// Individual Identity-02 commands
	identity0201Cmd := &cobra.Command{
		Use:   "identity-02-01",
		Short: "Check keystone.conf permissions",
		Run:   runIdentity0201Check,
	}

	identity0202Cmd := &cobra.Command{
		Use:   "identity-02-02",
		Short: "Check keystone-paste.ini permissions",
		Run:   runIdentity0202Check,
	}

	identity0203Cmd := &cobra.Command{
		Use:   "identity-02-03",
		Short: "Check policy.json permissions",
		Run:   runIdentity0203Check,
	}

	identity0204Cmd := &cobra.Command{
		Use:   "identity-02-04",
		Short: "Check logging.conf permissions",
		Run:   runIdentity0204Check,
	}

	identity0205Cmd := &cobra.Command{
		Use:   "identity-02-05",
		Short: "Check signing_cert.pem permissions",
		Run:   runIdentity0205Check,
	}

	identity0206Cmd := &cobra.Command{
		Use:   "identity-02-06",
		Short: "Check signing_key.pem permissions",
		Run:   runIdentity0206Check,
	}

	identity0207Cmd := &cobra.Command{
		Use:   "identity-02-07",
		Short: "Check ca.pem permissions",
		Run:   runIdentity0207Check,
	}

	identity0208Cmd := &cobra.Command{
		Use:   "identity-02-08",
		Short: "Check /etc/keystone directory permissions",
		Run:   runIdentity0208Check,
	}

	rootCmd.AddCommand(identity01Cmd)
	rootCmd.AddCommand(identity0101Cmd)
	rootCmd.AddCommand(identity0102Cmd)
	rootCmd.AddCommand(identity0103Cmd)
	rootCmd.AddCommand(identity0104Cmd)
	rootCmd.AddCommand(identity0105Cmd)
	rootCmd.AddCommand(identity0106Cmd)
	rootCmd.AddCommand(identity0107Cmd)
	rootCmd.AddCommand(identity0108Cmd)
	rootCmd.AddCommand(identity02Cmd)
	rootCmd.AddCommand(identity0201Cmd)
	rootCmd.AddCommand(identity0202Cmd)
	rootCmd.AddCommand(identity0203Cmd)
	rootCmd.AddCommand(identity0204Cmd)
	rootCmd.AddCommand(identity0205Cmd)
	rootCmd.AddCommand(identity0206Cmd)
	rootCmd.AddCommand(identity0207Cmd)
	rootCmd.AddCommand(identity0208Cmd)
}

func getSSHClient() (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: os.Getenv("SSH_USER"),
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("SSH_PASSWORD")),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return ssh.Dial("tcp", os.Getenv("SSH_HOST"), config)
}

func prettyPrintResult(result CheckResult) {
	fmt.Printf("\nCheck Result:\n")
	fmt.Printf("Status: %s\n", result.Status)
	fmt.Printf("Description: %s\n", result.Description)
	fmt.Printf("Details: %s\n", result.Details)
}

func runIdentity0101Check(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := CheckIdentity0101(client)
	prettyPrintResult(result)
}

func runIdentity0102Check(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := CheckIdentity0102(client)
	prettyPrintResult(result)
}

func runIdentity0103Check(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := CheckIdentity0103(client)
	prettyPrintResult(result)
}

func runIdentity0104Check(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := CheckIdentity0104(client)
	prettyPrintResult(result)
}

func runIdentity0105Check(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := CheckIdentity0105(client)
	prettyPrintResult(result)
}

func runIdentity0106Check(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := CheckIdentity0106(client)
	prettyPrintResult(result)
}

func runIdentity0107Check(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := CheckIdentity0107(client)
	prettyPrintResult(result)
}

func runIdentity0108Check(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := CheckIdentity0108(client)
	prettyPrintResult(result)
}

func runAllIdentity01Checks(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	checks := []struct {
		name string
		fn   func(*ssh.Client) CheckResult
	}{
		{"Identity-01-01", CheckIdentity0101},
		{"Identity-01-02", CheckIdentity0102},
		{"Identity-01-03", CheckIdentity0103},
		{"Identity-01-04", CheckIdentity0104},
		{"Identity-01-05", CheckIdentity0105},
		{"Identity-01-06", CheckIdentity0106},
		{"Identity-01-07", CheckIdentity0107},
		{"Identity-01-08", CheckIdentity0108},
	}

	for _, check := range checks {
		fmt.Printf("\n=== Running %s ===\n", check.name)
		result := check.fn(client)
		prettyPrintResult(result)
	}
}

// Add individual check functions for Identity-02
func runIdentity0201Check(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := CheckIdentity0201(client)
	prettyPrintResult(result)
}

func runIdentity0202Check(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := CheckIdentity0202(client)
	prettyPrintResult(result)
}

func runIdentity0203Check(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := CheckIdentity0203(client)
	prettyPrintResult(result)
}

func runIdentity0204Check(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := CheckIdentity0204(client)
	prettyPrintResult(result)
}

func runIdentity0205Check(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := CheckIdentity0205(client)
	prettyPrintResult(result)
}

func runIdentity0206Check(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := CheckIdentity0206(client)
	prettyPrintResult(result)
}

func runIdentity0207Check(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := CheckIdentity0207(client)
	prettyPrintResult(result)
}

func runIdentity0208Check(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := CheckIdentity0208(client)
	prettyPrintResult(result)
}

func runAllIdentity02Checks(cmd *cobra.Command, args []string) {
	client, err := getSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	checks := []struct {
		name string
		fn   func(*ssh.Client) CheckResult
	}{
		{"Identity-02-01", CheckIdentity0201},
		{"Identity-02-02", CheckIdentity0202},
		{"Identity-02-03", CheckIdentity0203},
		{"Identity-02-04", CheckIdentity0204},
		{"Identity-02-05", CheckIdentity0205},
		{"Identity-02-06", CheckIdentity0206},
		{"Identity-02-07", CheckIdentity0207},
		{"Identity-02-08", CheckIdentity0208},
	}

	for _, check := range checks {
		fmt.Printf("\n=== Running %s ===\n", check.name)
		result := check.fn(client)
		prettyPrintResult(result)
	}
}
