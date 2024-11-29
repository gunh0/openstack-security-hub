package cmd

import (
	"fmt"

	"github.com/gunh0/openstack-security-hub/checklist/keymanager"
	"github.com/gunh0/openstack-security-hub/util"
	"github.com/spf13/cobra"
)

func initKeyManagerCommands() {
	keyManager0101Cmd := &cobra.Command{
		Use:   "key-manager-01-01",
		Short: "Is user/group ownership of /etc/barbican/barbican.conf set to root:barbican?",
		Run:   runKeyManager0101Checks,
	}

	keyManager03Cmd := &cobra.Command{
		Use:   "key-manager-03",
		Short: "Is OpenStack Identity used for authentication?",
		Run:   runKeyManager03Checks,
	}

	RootCmd.AddCommand(keyManager0101Cmd)
	RootCmd.AddCommand(keyManager03Cmd)
}

func runKeyManager0101Checks(cmd *cobra.Command, args []string) {
	// Get SSH client
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	// Run check and print result
	result := keymanager.CheckKeyManager0101(client)
	util.PrettyPrintResult(result)
}

func runKeyManager03Checks(cmd *cobra.Command, args []string) {
	// Get SSH client
	client, err := util.GetSSHClient()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	// Run check and print result
	result := keymanager.CheckKeyManager03(client)
	util.PrettyPrintResult(result)
}
