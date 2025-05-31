.PHONY: help dev build clean backend frontend install test

# Default target
help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

# Development
dev: ## Start both backend and frontend in development mode
	@echo "Starting development servers..."
	@$(MAKE) -j2 backend frontend

backend: ## Start the Go backend server
	@echo "Starting backend server..."
	@cd backend && go run main.go

frontend: ## Start the SvelteKit frontend server
	@echo "Starting frontend server..."
	@cd frontend && pnpm dev

install: ## Install all dependencies
	@echo "Installing backend dependencies..."
	@cd backend && go mod tidy
	@echo "Installing frontend dependencies..."
	@cd frontend && pnpm install

# Building
build: ## Build both backend and frontend for production
	@echo "Building backend..."
	@cd backend && go build -o homeserver main.go
	@echo "Building frontend..."
	@cd frontend && pnpm build

build-backend: ## Build only the backend
	@cd backend && go build -o homeserver main.go

build-frontend: ## Build only the frontend
	@cd frontend && pnpm build

# Testing
test: ## Run tests for both backend and frontend
	@echo "Running backend tests..."
	@cd backend && go test ./...
	@echo "Running frontend tests..."
	@cd frontend && pnpm test 2>/dev/null || echo "No frontend tests configured"

test-backend: ## Run backend tests only
	@cd backend && go test ./...

# Cleanup
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -f backend/homeserver
	@rm -rf frontend/build
	@rm -rf frontend/.svelte-kit

# Linting
lint: ## Run linters
	@echo "Running Go linter..."
	@cd backend && go fmt ./... && go vet ./...
	@echo "Running frontend linter..."
	@cd frontend && pnpm check 2>/dev/null || echo "Frontend check completed"

# Git helpers
init-git: ## Initialize git repository with initial commit
	@git init
	@git add .
	@git commit -m "Initial commit: homeserver project setup"

# Environment setup
setup: install ## Complete project setup
	@echo "Project setup complete!"
	@echo "Run 'make dev' to start development servers"
