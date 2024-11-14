// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gunh0/openstack-security-hub/checklist"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
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

var identityCmd = &cobra.Command{
	Use:   "identity-01",
	Short: "Run identity-01 check",
	Run: func(cmd *cobra.Command, args []string) {
		runIdentityCheck()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(identityCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func startServer() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.GET("/health", healthCheck)
		api.GET("/check/identity-01", checkIdentity01)
	}

	r.Run(":8080")
}

func runIdentityCheck() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &ssh.ClientConfig{
		User: os.Getenv("SSH_USER"),
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("SSH_PASSWORD")),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", os.Getenv("SSH_HOST"), config)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer client.Close()

	result := checklist.CheckIdentity01(client)
	prettyPrintResult(result)
}

func prettyPrintResult(result checklist.CheckResult) {
	fmt.Printf("\nCheck Result:\n")
	fmt.Printf("Status: %s\n", result.Status)
	fmt.Printf("Description: %s\n", result.Description)
	fmt.Printf("Details: %s\n", result.Details)
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "OpenStack Security Hub is running",
	})
}

func checkIdentity01(c *gin.Context) {
	config := &ssh.ClientConfig{
		User: os.Getenv("SSH_USER"),
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("SSH_PASSWORD")),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", os.Getenv("SSH_HOST"), config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to connect to server",
			"error":   err.Error(),
		})
		return
	}
	defer client.Close()

	result := checklist.CheckIdentity01(client)
	c.JSON(http.StatusOK, result)
}
