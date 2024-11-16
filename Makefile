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

# API endpoints
API_BASE=http://localhost:8080/api/v1

.PHONY: build clean run dev test-health api-check cli-check

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE)

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf ./tmp

# Run the server
run:
	$(GORUN) $(MAIN_FILE) server

# Run with air (hot reload)
dev:
	$(GOAIR)

# Test health endpoint
test-health:
	curl -X GET $(API_BASE)/health

# API checks
api-identity-01:
	curl -s $(API_BASE)/check/identity-01 | python3 -m json.tool

api-identity-01-01:
	curl -s $(API_BASE)/check/identity-01-01 | python3 -m json.tool

api-identity-01-02:
	curl -s $(API_BASE)/check/identity-01-02 | python3 -m json.tool

api-identity-01-03:
	curl -s $(API_BASE)/check/identity-01-03 | python3 -m json.tool

api-identity-01-04:
	curl -s $(API_BASE)/check/identity-01-04 | python3 -m json.tool

api-identity-01-05:
	curl -s $(API_BASE)/check/identity-01-05 | python3 -m json.tool

api-identity-01-06:
	curl -s $(API_BASE)/check/identity-01-06 | python3 -m json.tool

api-identity-01-07:
	curl -s $(API_BASE)/check/identity-01-07 | python3 -m json.tool

api-identity-01-08:
	curl -s $(API_BASE)/check/identity-01-08 | python3 -m json.tool

# CLI checks
cli-identity-01:
	$(GORUN) $(MAIN_FILE) identity-01

cli-identity-01-01:
	$(GORUN) $(MAIN_FILE) identity-01-01

cli-identity-01-02:
	$(GORUN) $(MAIN_FILE) identity-01-02

cli-identity-01-03:
	$(GORUN) $(MAIN_FILE) identity-01-03

cli-identity-01-04:
	$(GORUN) $(MAIN_FILE) identity-01-04

cli-identity-01-05:
	$(GORUN) $(MAIN_FILE) identity-01-05

cli-identity-01-06:
	$(GORUN) $(MAIN_FILE) identity-01-06

cli-identity-01-07:
	$(GORUN) $(MAIN_FILE) identity-01-07

cli-identity-01-08:
	$(GORUN) $(MAIN_FILE) identity-01-08

# Run all checks
check-all:
	$(GORUN) $(MAIN_FILE) identity-01

# Help
help:
	@echo "Available commands:"
	@echo "  make build              - Build the binary"
	@echo "  make clean              - Clean build files"
	@echo "  make run               - Run the server"
	@echo "  make dev               - Run with hot reload"
	@echo "  make test-health       - Test health endpoint"
	@echo "  make api-identity-01   - Run all identity checks via API"
	@echo "  make api-identity-01-XX - Run specific identity check via API"
	@echo "  make cli-identity-01   - Run all identity checks via CLI"
	@echo "  make cli-identity-01-XX - Run specific identity check via CLI"
	@echo "  make check-all         - Run all checks"