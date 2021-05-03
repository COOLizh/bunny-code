PROJECT_NAME := "task_repo"
APP_NAME := "app"
PKG := "./"
CMD := "$(PKG)/cmd/$(APP_NAME)"

PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all build test lint get_lint get_dep setup_env containerise run

all: build

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}
	@golangci-lint run --skip-dirs='pkg/pb'

test: ## Run unittests
	@go test -v -race ${PKG_LIST}

get_dep: ## Get the dependencies
	@go get -v -t -d ./...
	@go get -u golang.org/x/lint/golint
	@go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.31.0

build: ## Build the binary file
	@go build -i -v $(CMD)

setup_env:
	@cp ./env.example ./.env

containerise: ## Build Docker image
	@docker build . -t taskapp

run: ## Run Docker container
	@docker run --env-file .env -p 8808:8808 taskapp

proto: ## Generate the gRPC client and server interfaces from .proto service definition
	@protoc -I . ./proto/*.proto \
    --go_out=. \
    --go-grpc_out=require_unimplemented_servers=false:.