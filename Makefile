.PHONY: run build test docker-run docker_down

all: build test

# TODO - Add more commands here to run tests, build, etc.
run:
	@go run cmd/api/main.go

build:
	@go build -o bin/svr cmd/svr/main.go

test:
	@echo "Running go tests..."
	@go test -v ./...

docker-run:
	@docker-compose up --build

docker-down:
	@docker-compose down