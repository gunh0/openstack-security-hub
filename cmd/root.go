// cmd/root.go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "security-hub",
	Short: "OpenStack Security Hub CLI",
	Long: `A CLI tool for OpenStack security checking.
Run 'security-hub help' for usage information.`,
	// No arguments starts the server
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Starting server... Use Ctrl+C to stop")
		}
	},
}

func init() {
	// Initialize all service commands
	initIdentityCommands()
	initDashboardCommands()
	initKeyManagerCommands()

	// Add help command
	helpCmd := &cobra.Command{
		Use:   "help",
		Short: "Show help for all commands",
		Run: func(cmd *cobra.Command, args []string) {
			RootCmd.Help()
		},
	}
	RootCmd.AddCommand(helpCmd)
}
