.PHONY: all
all: run

.PHONY: run
run:
	@go run cmd/main.go

.PHONY: test
test:
	@go test ./...
