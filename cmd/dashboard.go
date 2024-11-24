package cmd

import (
	"fmt"

	"github.com/gunh0/openstack-security-hub/checklist/dashboard"
	"github.com/gunh0/openstack-security-hub/util"
	"github.com/spf13/cobra"
)

func initDashboardCommands() {
	dashboard01Cmd := &cobra.Command{
		Use:   "dashboard-01",
		Short: "Is user/group of config files set to root/horizon?",
		Run:   runDashboard01Checks,
	}

	dashboard04Cmd := &cobra.Command{
		Use:   "dashboard-04",
		Short: "Is CSRF_COOKIE_SECURE parameter set to True?",
		Run:   runDashboard04Checks,
	}

	dashboard05Cmd := &cobra.Command{
		Use:   "dashboard-05",
		Short: "Is SESSION_COOKIE_SECURE parameter set to True?",
		Run:   runDashboard05Checks,
	}

	RootCmd.AddCommand(dashboard01Cmd)
	RootCmd.AddCommand(dashboard04Cmd)
	RootCmd.AddCommand(dashboard05Cmd)
}

func runDashboard01Checks(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := dashboard.CheckDashboard01(client)
	util.PrettyPrintResult(result)
}

func runDashboard04Checks(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := dashboard.CheckDashboard04(client)
	util.PrettyPrintResult(result)
}

func runDashboard05Checks(cmd *cobra.Command, args []string) {
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	result := dashboard.CheckDashboard05(client)
	util.PrettyPrintResult(result)
}
