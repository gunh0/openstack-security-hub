// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gunh0/openstack-security-hub/api"
	"github.com/gunh0/openstack-security-hub/cmd"
	"github.com/gunh0/openstack-security-hub/docs"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           OpenStack Security Hub API
// @version         1.0
// @description     API server for OpenStack security checking
// @BasePath        /api/v1

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// If no arguments, start server
	if len(os.Args) == 1 {
		startServer()
		return
	}

	// Otherwise, execute the command
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func startServer() {
	r := gin.Default()

	docs.SwaggerInfo.Title = "OpenStack Security Hub API"
	docs.SwaggerInfo.Description = "API server for OpenStack security checking"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Swagger initialization
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register all API routes and health check endpoint
	api.RegisterRoutes(r)

	log.Printf("Starting server on :8080...")
	log.Printf("Swagger documentation available at http://localhost:8080/swagger/index.html")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
