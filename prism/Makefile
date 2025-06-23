# Prism Framework Makefile

.PHONY: help build test test-unit test-e2e run-demo clean lint fmt

# Default target
help:
	@echo "Prism Framework Build Commands:"
	@echo ""
	@echo "  build      - Build the framework and demo app"
	@echo "  test       - Run all tests (unit + e2e)"
	@echo "  test-unit  - Run unit tests"
	@echo "  test-e2e   - Run E2E tests"
	@echo "  run-demo   - Start the demo server"
	@echo "  clean      - Clean build artifacts"
	@echo "  lint       - Run linter"
	@echo "  fmt        - Format code"
	@echo ""

# Build targets
build:
	@echo "Building prism framework..."
	go build .
	@echo "Building demo app..."
	cd demo && go build -o demo-server .
	@echo "Build complete!"

# Test targets
test: test-unit test-e2e

test-unit:
	@echo "Running unit tests..."
	go test -v -race ./...

test-e2e:
	@echo "Running E2E tests..."
	@echo "Note: This will start a demo server and run browser tests"
	cd tests/e2e && go test -v -timeout 5m .

# Development targets
run-demo:
	@echo "Starting demo server on http://localhost:9090"
	@echo "Available routes:"
	@echo "  / - Home page"
	@echo "  /login - Login page"
	@echo "  /demo - Interactive demo"
	@echo ""
	cd demo && go run .

run-demo-build:
	@echo "Starting demo server (built binary)..."
	cd demo && ./demo-server

# Maintenance targets
clean:
	@echo "Cleaning build artifacts..."
	rm -f demo/demo-server
	go clean
	@echo "Clean complete!"

lint:
	@echo "Running linter..."
	golangci-lint run ./...

fmt:
	@echo "Formatting code..."
	go fmt ./...
	@echo "Format complete!"

# Development helpers
dev-setup:
	@echo "Setting up development environment..."
	go mod tidy
	@echo "Installing golangci-lint..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Development setup complete!"

# Quick test cycle
quick-test: fmt build test-unit
	@echo "Quick test cycle complete!"

# Full test cycle
full-test: fmt build test
	@echo "Full test cycle complete!"

# Watch mode (requires entr)
watch:
	@echo "Watching for changes..."
	find . -name "*.go" | entr -r make quick-test

# Release build
release: clean fmt lint test build
	@echo "Release build complete!"
	@echo "Demo server binary: demo/demo-server"

# Help for specific commands
help-e2e:
	@echo "E2E Testing Guide:"
	@echo ""
	@echo "The E2E tests use Rod (Chrome browser automation) to test the framework."
	@echo ""
	@echo "Prerequisites:"
	@echo "  - Chrome/Chromium browser installed"
	@echo "  - No other services running on port 9090"
	@echo ""
	@echo "What the tests do:"
	@echo "  1. Build the demo app"
	@echo "  2. Start demo server on port 9090"
	@echo "  3. Launch headless Chrome browser"
	@echo "  4. Run automated tests against the running app"
	@echo "  5. Clean up server and browser"
	@echo ""
	@echo "Test categories:"
	@echo "  - Basic page loading and navigation"
	@echo "  - Modal and toast interactions"
	@echo "  - Form validation and user input"
	@echo "  - Data table operations"
	@echo "  - Responsive design"
	@echo "  - Accessibility compliance"
	@echo ""
	@echo "Run with: make test-e2e"

help-demo:
	@echo "Demo Application Guide:"
	@echo ""
	@echo "The demo app showcases all framework components:"
	@echo ""
	@echo "Routes:"
	@echo "  /      - Basic home page"
	@echo "  /login - Login form (auth not implemented yet)"
	@echo "  /demo  - Interactive component showcase"
	@echo ""
	@echo "Demo page features:"
	@echo "  - Toast notifications (info, success, warning, error)"
	@echo "  - Modal dialogs (basic and confirm dialogs)"
	@echo "  - Data table with selection and actions"
	@echo "  - State management demonstrations"
	@echo ""
	@echo "Start with: make run-demo"
	@echo "Then visit: http://localhost:9090/demo"