TAGS="dev"

# Prints a list of useful targets.
help:
	@echo ""
	@echo "charmil CLI"
	@echo ""
	@echo "make lint                 	run golangci-lint"
	@echo "make test                 	run  tests"
	@echo "make format             		format files"
	@echo "make build             		build files"

.PHONY: help

run:
	go run ./examples/host

build:
	go build -o host_example ./examples/host
	go build -o validator_example ./validator/example

build/tags:
	go build -tags ${TAGS} ./examples/host

test/validator:
	go test ./validator/example

test/unit: install
	go test ./core/...
.PHONY: test/unit
	
	
# Requires golangci-lint to be installed @ $(go env GOPATH)/bin/golangci-lint
# https://golangci-lint.run/usage/install/
lint:
	golangci-lint run cmd/... pkg/... internal/...
.PHONY: lint


# clean up code and dependencies
format:
	@go mod tidy

	@gofmt -w `find . -type f -name '*.go'`
.PHONY: format
