# Simple Makefile for a Go project

# Build the application
all: build

build: clean
	@go build -o main cmd/api/main.go

setup:
	@go get tool

# Run the application
run:
	@go run cmd/api/main.go

goose:
	@read -p "Action: " action; \
	@go tool goose -dir ./db/migrations postgres "user=postgres password=password host=localhost dbname=gobp sslmode=disable" $$action

migrate:
	@go tool goose -dir ./db/migrations postgres "user=postgres password=password host=localhost dbname=gobp sslmode=disable" up

create-migration:
	@read -p "Enter migration name: " name; \
	@go tool goose -dir ./db/migrations create $$name sql

# Live Reload
watch:
	@echo "Watching..."
	@go tool air

query:
	@go tool sqlc generate

.PHONY: all setup build run clean migrate create-migration watch
