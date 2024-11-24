.PHONY: run build test docker-run docker_down watch lint format integration-test

all: build test

run:
	@go run cmd/api/main.go

build:
	@go build -o bin/svr cmd/svr/main.go

test:
	@echo "Running go tests..."
	@go test -v ./...

docker-run:
	@docker-compose up --build

watch:
	@air -c .air.toml

docker-down:
	@docker-compose down

lint:
	@golangci-lint run

format:
	@gofmt -s -w .

integration-test:
	@go test -tags=integration -v ./...