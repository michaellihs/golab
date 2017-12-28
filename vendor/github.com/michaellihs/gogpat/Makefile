PROJECT_NAME := "gogpat"
PKG := "github.com/solidnerd/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep build clean test coverage coverhtml lint ci

all: build

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}

ci: ## Runs test in an ci bootstaps a gitlab instance 
	ci/test.sh ${PKG_LIST}

race: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}
	
dep: ## Get the dependencies
	@go get -v -d ./...
	@go get -u github.com/golang/lint/golint

build: dep ## Build the binary file
	@go build -i -o bin/${PROJECT_NAME} -v $(PKG)

clean: ## Remove previous build
	@rm -rf bin/

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}
