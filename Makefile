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

.PHONY: build clean run dev test-health api-check cli-check swagger-init swagger-fmt

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

cli-keymanager-%:
	$(GORUN) $(MAIN_FILE) keymanager-$*

# Run all checks via CLI
cli-check-all:
	@echo "Running all Identity checks..."
	$(MAKE) cli-identity-01
	$(MAKE) cli-identity-02
	$(MAKE) cli-identity-03
	$(MAKE) cli-identity-05
	@echo "Running all Dashboard checks..."
	$(MAKE) cli-dashboard-01
	$(MAKE) cli-dashboard-02
	$(MAKE) cli-dashboard-03
	$(MAKE) cli-dashboard-04
	$(MAKE) cli-dashboard-05
	@echo "Running all Keymanager checks..."
	$(MAKE) cli-keymanager-01

# API checks
api-identity-%:
	curl -s $(API_BASE)/check/identity-$* | python3 -m json.tool

api-dashboard-%:
	curl -s $(API_BASE)/check/dashboard-$* | python3 -m json.tool

api-keymanager-%:
	curl -s $(API_BASE)/check/keymanager-$* | python3 -m json.tool

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
	@echo "  make swagger-fmt          - Format Swagger documentation"