# Simple Makefile for a Go project

# Build the application
all: build

build: clean
	@go build -o main cmd/api/main.go

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
	@go test ./... -v

# Clean the binary
clean:
	@rm -f main

goose:
	@read -p "Action: " action; \
	goose -dir ./db/migrations postgres "user=postgres password=password dbname=gobp sslmode=disable" $$action

migrate:
	@goose -dir ./db/migrations postgres "user=postgres password=password dbname=gobp sslmode=disable" up

create-migration:
	@read -p "Enter migration name: " name; \
	goose -dir ./db/migrations create $$name sql

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean
