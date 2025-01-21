# Simple Makefile for a Go project

# Build the application
all: build

build: clean
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

setup:
	@read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/air-verse/air@latest; \
			echo "Watching...";\
	else \
			echo "You chose not to install air."; \
	fi
	@read -p "Go's 'goose' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/pressly/goose/v3/cmd/goose@latest; \
	else \
			echo "You chose not to install goose."; \
	fi
	@read -p "Go's 'sqlc' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.27.0 \
	else \
			echo "You chose not to install sqlc."; \
	fi

goose:
	@read -p "Action: " action; \
	@goose -dir ./db/migrations postgres "user=postgres password=password dbname=gobp sslmode=disable" $$action

migrate:
	@goose -dir ./db/migrations postgres "user=postgres password=password dbname=gobp sslmode=disable" up

create-migration:
	@read -p "Enter migration name: " name; \
	@goose -dir ./db/migrations create $$name sql

# Live Reload
watch:
	@echo "Watching..."
	@air

.PHONY: all setup build run clean migrate create-migration watch
