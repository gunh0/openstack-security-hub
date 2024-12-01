# Binary name
BINARY_NAME=openstack-security-hub

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOAIR=air
SWAG=swag

# Main entry point
MAIN_FILE=main.go

# API endpoints
API_BASE=http://localhost:8080/api/v1

.PHONY: build clean run dev test-health api-check cli-check swagger-init

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE)

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf ./tmp

# Generate Swagger documentation
swagger-init:
	$(SWAG) init

# Run the server (no arguments for server mode)
run: swagger-init
	go mod tidy
	$(GORUN) $(MAIN_FILE)

# Run with air (hot reload)
dev: swagger-init
	go mod tidy
	$(GOAIR)

# Test health endpoint
test-health:
	curl -X GET $(API_BASE)/health

# CLI commands
cli-identity-%:
	$(GORUN) $(MAIN_FILE) identity-$*

cli-dashboard-%:
	$(GORUN) $(MAIN_FILE) dashboard-$*

cli-key-manager-%:
	$(GORUN) $(MAIN_FILE) key-manager-$*

# Run all checks via CLI
cli-check-all:
	@echo "Running all Identity checks..."
	$(MAKE) cli-identity-01-01
	@echo "Running all Dashboard checks..."
	$(MAKE) cli-dashboard-04
	$(MAKE) cli-dashboard-05
	$(MAKE) cli-dashboard-06
	@echo "Running all Secrets Management checks..."
	$(MAKE) cli-key-manager-01-01

# API checks
api-identity-%:
	curl -s $(API_BASE)/check/identity-$* | python3 -m json.tool

api-dashboard-%:
	curl -s $(API_BASE)/check/dashboard-$* | python3 -m json.tool

api-key-manager-%:
	curl -s $(API_BASE)/check/key-manager-$* | python3 -m json.tool

# Help
help:
	@echo "Available commands:"
	@echo "Server commands:"
	@echo "  make run                  - Run the server"
	@echo "  make dev                  - Run with hot reload"
	@echo "  make test-health          - Test health endpoint"
	@echo ""
	@echo "CLI commands:"
	@echo "  make cli-identity-XX      - Run specific identity check"
	@echo "  make cli-dashboard-XX     - Run specific dashboard check"
	@echo "  make cli-keymanager-XX    - Run specific keymanager check"
	@echo "  make cli-check-all        - Run all checks"
	@echo ""
	@echo "API commands:"
	@echo "  make api-identity-XX      - Run specific identity check via API"
	@echo "  make api-dashboard-XX     - Run specific dashboard check via API"
	@echo "  make api-keymanager-XX    - Run specific keymanager check via API"
	@echo ""
	@echo "Development commands:"
	@echo "  make build                - Build the binary"
	@echo "  make clean                - Clean build files"
	@echo "  make swagger-init         - Generate Swagger documentation"