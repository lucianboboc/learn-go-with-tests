.PHONY: lint run test

run:
	go run hello.go

lint:
	golangci-lint run

test:
	go test ./...