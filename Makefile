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
	@echo "make setup/git/hooks      	setup git hooks"
.PHONY: help

run:
	go run ./examples/host

build:
	go build -o host_example ./examples/host

build/tags:
	go build -tags ${TAGS} ./examples/host

test/validator:
	go test ./validator/rules

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

# Set git hook path to .githooks/
setup/githooks:
	git config core.hooksPath .githooks
.PHONY: setup/git/hooks
