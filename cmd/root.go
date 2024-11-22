package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "security-hub",
	Short: "OpenStack Security Hub CLI",
	Long:  `A CLI tool for OpenStack security checking.`,
}

func init() {
	// Initialize all service commands
	initIdentityCommands()
	initDashboardCommands()
	initKeyManagerCommands()
}
