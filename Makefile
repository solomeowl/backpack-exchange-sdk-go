.PHONY: all build test lint clean fmt vet tidy

all: lint test build

build:
	go build ./...

test:
	go test -v -race -cover ./...

lint:
	golangci-lint run ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	go clean
	rm -rf bin/

# Run examples
run-public:
	go run ./examples/public/main.go

run-authenticated:
	go run ./examples/authenticated/main.go

run-websocket:
	go run ./examples/websocket/main.go
