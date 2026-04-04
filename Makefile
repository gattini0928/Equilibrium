APP_NAME=jwt-auth

build:
	@go build -o bin/$(APP_NAME) cmd/server/main.go

run: build
	@./bin/$(APP_NAME)

dev:
	go run cmd/server/main.go

test:
	@go test -v ./...

migrate:
	go run cmd/migrate/main.go