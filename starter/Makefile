.DEFAULT_GOAL := help
SHELL = bash

# see internal/build.go on build configurations
CLI_VERSION ?= "dev"

GO_LDFLAGS := -X github.com/aerogear/charmil/internal/build.Version=$(CLI_VERSION) $(GO_LDFLAGS)
BUILDFLAGS :=

ifdef DEBUG
BUILDFLAGS := -gcflags "all=-N -l" $(BUILDFLAGS)
endif

# The details of the application:
binary:=placeholdercli

# Enable Go modules:
export GO111MODULE=on

# Prints a list of useful targets.
help:
	@echo ""
	@echo "starter cli"
	@echo ""
	@echo "make lint                 	run golangci-lint"
	@echo "make binary               	compile binaries"
	@echo "make test                 	run  tests"
	@echo "make format             		format files"
	@echo "make docs/check						check if docs need to be updated"
	@echo "make docs/generate					generate the docs"

	@echo "$(fake)"
.PHONY: help

# Requires golangci-lint to be installed @ $(go env GOPATH)/bin/golangci-lint
# https://golangci-lint.run/usage/install/
lint:
	golangci-lint run cmd/... pkg/... internal/...
.PHONY: lint

generate:
	go generate ./...

# Build binaries
# NOTE it may be necessary to use CGO_ENABLED=0 for backwards compatibility with centos7 if not using centos7
binary:
	go build $(BUILDFLAGS) -ldflags "${GO_LDFLAGS}" -o ${binary} ./cmd/${binary}
.PHONY: binary

install:
	go install -trimpath $(BUILDFLAGS) -ldflags "${GO_LDFLAGS}" ./cmd/${binary}
.PHONY: install

test/unit: install
	go test ./cmd/... ./pkg/...
.PHONY: test/unit

# clean up code and dependencies
format:
	@go mod tidy

	@gofmt -w `find . -type f -name '*.go'`
.PHONY: format

# Symlink common git hookd into .git directory
githooks:
	ln -fs $$(pwd)/githooks/pre-commit .git/hooks
.PHONY: githooks

docs/check: docs/generate
	./scripts/check-docs.sh
.PHONY: docs/check

docs/generate:
	GENERATE_DOCS=true go run ./cmd/${binary}
.PHONY: docs/generate

docs/generate-modular-docs: docs/generate
	SRC_DIR=$$(pwd)/docs/commands DEST_DIR=$$(pwd)/dist go run ./cmd/modular-docs
.PHONY: docs/generate-modular-docs
