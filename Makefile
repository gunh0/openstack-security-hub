# Binary name
BINARY_NAME=openstack-security-hub

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOAIR=air

# Main entry point
MAIN_FILE=main.go

.PHONY: build clean run dev test-health check-identity-01 cli-identity-01

# Run the application
run:
	$(GORUN) $(MAIN_FILE) server

# Run with air (hot reload)
dev:
	$(GOAIR)

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE)

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf ./tmp

# Test health endpoint
test-health:
	curl -X GET http://localhost:8080/api/v1/health

# Check identity-01 (API)
api-identity-01:
	curl -s http://localhost:8080/api/v1/check/identity-01 | python3 -m json.tool

# Check identity-01 (CLI)
cli-identity-01:
	$(GORUN) $(MAIN_FILE) identity-01