.PHONY: build run

build:
	@go build -o bin/api ./cmd/api

run: build
	@./bin/api

# database migrations
migrate-up:
	@go run ./cmd/migrate up

migrate-down:
	@go run ./cmd/migrate down
