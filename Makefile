# Binary name
BINARY_NAME=openstack-security-hub

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean

# Main entry point
MAIN_FILE=main.go

.PHONY: build clean run test-health

# Run the application
run:
	$(GORUN) $(MAIN_FILE)

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE)

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Test health endpoint
test-health:
	curl -X GET http://localhost:8080/api/v1/health