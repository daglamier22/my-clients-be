# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	
	
	@go build -o ./bin/main cmd/api/*.go

# Run the application
run:
	@go run cmd/api/main.go

# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
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
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/air-verse/air@latest; \
	        air; \
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

# Goose migration status
goose-db-status:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable&search_path=$(DB_SCHEMA)" goose -dir="./internal/database/migrations" status

# Goose migration status
goose-up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable&search_path=$(DB_SCHEMA)" goose -dir="./internal/database/migrations" up

# Goose migration status
goose-reset:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable&search_path=$(DB_SCHEMA)" goose -dir="./internal/database/migrations" reset

.PHONY: all build run docker-run docker-down test clean watch goose-db-status goose-up goose-reset
