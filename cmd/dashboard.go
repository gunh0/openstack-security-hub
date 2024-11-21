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

	RootCmd.AddCommand(dashboard01Cmd)
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
