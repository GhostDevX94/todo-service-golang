ifneq (,$(wildcard .env))
    include .env
else
    ifneq (,$(wildcard env.example))
        $(shell cp env.example .env)
        include .env
    endif
endif

MIGRATE ?= /opt/homebrew/bin/migrate


build:
	@echo "Putting the application together..."
	go build -o $(BINARY_PATH) cmd/main.go

clean:
	@echo "Removing binary file..."
	rm -f $(BINARY_PATH)

run: build
	@echo "Starting $(BINARY_PATH) on port $(APP_PORT)..."
	@$(BINARY_PATH) &

stop:
	@echo "Stopping $(BINARY_NAME)..."
	@pkill -f "$(BINARY_PATH)" || true

restart: stop build run
	@echo "Restart completed."
	
swagger:
	@echo "Generating Swagger documentation..."
	@~/go/bin/swag init -g cmd/main.go -o docs --dir cmd,internal/http,internal/dto,internal/model
	@echo "Swagger docs generated successfully!"

docker-build:
	@echo "Building Docker image..."
	docker build -t todo-list-api:latest .

docker-up:
	@echo "Starting services with Docker Compose..."
	docker compose up -d
	@echo "Services started! API available at http://localhost:8181"

docker-down:
	@echo "Stopping Docker Compose services..."
	docker compose down

docker-logs:
	@echo "Showing logs..."
	docker compose logs -f app

docker-rebuild:
	@echo "Rebuilding and restarting services..."
	docker compose down
	docker compose up -d --build

# Development helpers
lint:
	@echo "Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Install with: brew install golangci-lint"; \
	fi

fmt:
	@echo "Formatting code..."
	@go fmt ./...

test:
	@echo "Running tests..."
	@go test -v -cover ./...

.PHONY: build clean run stop restart db-up create-migration migrate-up migrate-down migrate-steps migrate-rollback swagger docker-build docker-up docker-down docker-logs docker-rebuild lint fmt test


