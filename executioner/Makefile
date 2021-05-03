PROJECT_NAME := "executioner"
APP_NAME := "app"
PKG := "./"
CMD := "$(PKG)/cmd/$(APP_NAME)"

PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all build test lint get_lint get_dep containerise set_env run

all: build

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}
	@golangci-lint run

test: ## Run unit tests
	@go test -v -race ${PKG_LIST}

get_dep: ## Get the dependencies
	@go get -v -t -d ./...
	@go get -u golang.org/x/lint/golint
	@go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.31.0

build: ## Build the binary file
	@go build -i -v $(CMD)

containerise: ## Build Docker image
	@docker build . -t executioner

set_env:
	@cp ./env.example ./.env

run: ## Run Docker container
	@docker run --env-file .env -p 8001:8001 executioner