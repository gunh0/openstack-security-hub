package cmd

import (
	"fmt"

	"github.com/gunh0/openstack-security-hub/checklist"
	"github.com/gunh0/openstack-security-hub/checklist/identity"
	"github.com/gunh0/openstack-security-hub/util"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

func initIdentityCommands() {
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

	identity03Cmd := &cobra.Command{
		Use:   "identity-03",
		Short: "Is TLS enabled for Identity?",
		Run:   runIdentity03Check,
	}

	identity05Cmd := &cobra.Command{
		Use:   "identity-05",
		Short: "Is max_request_body_size set to default (114688)?",
		Run:   runIdentity05Check,
	}

	identity06Cmd := &cobra.Command{
		Use:   "identity-06",
		Short: "Disable admin token in /etc/keystone/keystone.conf",
		Run:   runIdentity06Check,
	}

	RootCmd.AddCommand(identity01Cmd)
	RootCmd.AddCommand(identity0101Cmd)
	RootCmd.AddCommand(identity0102Cmd)
	RootCmd.AddCommand(identity0103Cmd)
	RootCmd.AddCommand(identity0104Cmd)
	RootCmd.AddCommand(identity0105Cmd)
	RootCmd.AddCommand(identity0106Cmd)
	RootCmd.AddCommand(identity0107Cmd)
	RootCmd.AddCommand(identity0108Cmd)
	RootCmd.AddCommand(identity02Cmd)
	RootCmd.AddCommand(identity0201Cmd)
	RootCmd.AddCommand(identity0202Cmd)
	RootCmd.AddCommand(identity0203Cmd)
	RootCmd.AddCommand(identity0204Cmd)
	RootCmd.AddCommand(identity0205Cmd)
	RootCmd.AddCommand(identity0206Cmd)
	RootCmd.AddCommand(identity0207Cmd)
	RootCmd.AddCommand(identity0208Cmd)
	RootCmd.AddCommand(identity03Cmd)
	RootCmd.AddCommand(identity05Cmd)
	RootCmd.AddCommand(identity06Cmd)
}

func runIdentity0101Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0101(client)
	util.PrettyPrintResult(result)
}

func runIdentity0102Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0102(client)
	util.PrettyPrintResult(result)
}

func runIdentity0103Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0103(client)
	util.PrettyPrintResult(result)
}

func runIdentity0104Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0104(client)
	util.PrettyPrintResult(result)
}

func runIdentity0105Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0105(client)
	util.PrettyPrintResult(result)
}

func runIdentity0106Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0106(client)
	util.PrettyPrintResult(result)
}

func runIdentity0107Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0107(client)
	util.PrettyPrintResult(result)
}

func runIdentity0108Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0108(client)
	util.PrettyPrintResult(result)
}

func runAllIdentity01Checks(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	checks := []struct {
		name string
		fn   func(*ssh.Client) checklist.CheckResult
	}{
		{"Identity-01-01", identity.CheckIdentity0101},
		{"Identity-01-02", identity.CheckIdentity0102},
		{"Identity-01-03", identity.CheckIdentity0103},
		{"Identity-01-04", identity.CheckIdentity0104},
		{"Identity-01-05", identity.CheckIdentity0105},
		{"Identity-01-06", identity.CheckIdentity0106},
		{"Identity-01-07", identity.CheckIdentity0107},
		{"Identity-01-08", identity.CheckIdentity0108},
	}

	for _, check := range checks {
		result := check.fn(client)
		util.PrettyPrintResult(result)
	}
}

// Add individual check functions for Identity-02
func runIdentity0201Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0201(client)
	util.PrettyPrintResult(result)
}

func runIdentity0202Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0202(client)
	util.PrettyPrintResult(result)
}

func runIdentity0203Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0203(client)
	util.PrettyPrintResult(result)
}

func runIdentity0204Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0204(client)
	util.PrettyPrintResult(result)
}

func runIdentity0205Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0205(client)
	util.PrettyPrintResult(result)
}

func runIdentity0206Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0206(client)
	util.PrettyPrintResult(result)
}

func runIdentity0207Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0207(client)
	util.PrettyPrintResult(result)
}

func runIdentity0208Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity0208(client)
	util.PrettyPrintResult(result)
}

func runAllIdentity02Checks(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	checks := []struct {
		name string
		fn   func(*ssh.Client) checklist.CheckResult
	}{
		{"Identity-02-01", identity.CheckIdentity0201},
		{"Identity-02-02", identity.CheckIdentity0202},
		{"Identity-02-03", identity.CheckIdentity0203},
		{"Identity-02-04", identity.CheckIdentity0204},
		{"Identity-02-05", identity.CheckIdentity0205},
		{"Identity-02-06", identity.CheckIdentity0206},
		{"Identity-02-07", identity.CheckIdentity0207},
		{"Identity-02-08", identity.CheckIdentity0208},
	}

	for _, check := range checks {
		result := check.fn(client)
		util.PrettyPrintResult(result)
	}
}

func runIdentity03Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity03(client)
	util.PrettyPrintResult(result)
}

func runIdentity05Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity05(client)
	util.PrettyPrintResult(result)
}

func runIdentity06Check(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := identity.CheckIdentity06(client)
	util.PrettyPrintResult(result)
}
