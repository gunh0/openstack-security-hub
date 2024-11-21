package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gunh0/openstack-security-hub/api"
	"github.com/gunh0/openstack-security-hub/cmd"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "security-hub",
	Short: "OpenStack Security Hub CLI",
	Long:  `A CLI tool for OpenStack security checking.`,
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the API server",
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func init() {
	cmd.RootCmd.AddCommand(serverCmd)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func startServer() {
	r := gin.Default()

	// Register all routes including health check
	api.RegisterRoutes(r)

	log.Printf("Starting server on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
