# Simple Makefile for a Go project
# Build the application
all: build test

format:
	@echo "Formatting..."
	@go fmt ./...

build:
	@echo "Building..."
	
	
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go


docker-run-db:
	@docker-compose up dev-db-postgres

# Create DB container
docker-run:
	@if docker compose --env-file app.env up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose --env-file app.env up --build; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v -cover
# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Clear cache
clear-cache:
	@echo "Clearing cache..."
	@go clean -cache

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

sqlc:
	@sqlc generate

migrate-up:
	@migrate -path internal/database/migration \
			 -database "postgresql://user:password@localhost:5432/secure-real-estate-app?sslmode=disable" \
			 -verbose up

migrate-down:
	@migrate -path internal/database/migration \
			 -database "postgresql://user:password@localhost:5432/secure-real-estate-app?sslmode=disable" \
			 -verbose down

mock:
	@mockgen -destination internal/database/mock/service.go \
			 backend_real_estate/internal/database \
			 Service

swagger-gen:
	@swag init -g ./cmd/api/main.go -o ./docs/api

.PHONY: all build run test clean watch docker-run docker-down itest sqlc migrate-up migrate-down mock swagger-gen
