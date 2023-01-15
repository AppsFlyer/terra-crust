SHELL := /bin/bash
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

export GO_VERSION ?= $(shell grep -E "^go " go.mod | awk -F' ' '{print $$2}' )
export GOLANGCI_LINT_VERSION=v1.51.1
export MAIN_GO=internal/cmd/main.go

.PHONY: help
help: ## Show this help message.

	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


get-go-dependancies: 	## Get dependancies as devlared in the project 
	$(info)
	$(info ========[ $@ ]========)
	go get -v -t -d ./...

get-linter:		## Get golangci-lint
	$(info)
	$(info ========[ $@ ]========)
	@if ! golangci-lint --version 2>/dev/null; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin ${GOLANGCI_LINT_VERSION};\
	fi;\

lint: get-linter ## Run linter
	$(info)
	$(info ========[ $@ ]========)
	rm lint.xml
	golangci-lint run

test: ## Run go tests
	$(info)
	$(info ========[ $@ ]========)
	go test ./...	

build: get-go-dependancies lint ## Build
	$(info)
	$(info ========[ $@ ]========)
	go build -v ${MAIN_GO}


compile: build ## Compile binaries for multiple OSs and architectures
	$(info)
	$(info ========[ $@ ]========)
	@for os in "linux" "darwin"; do \
		for arch in "amd64" "arm64"; do \
			echo "$${os} - $${arch}"; \
			GOOS=$${os} GOARCH=$${arch} go build -o bin/terra-crust-$${os}-$${arch} ${MAIN_GO}; \
		done \
	done

build-docker: ## Build a docker image
	$(info)
	$(info ========[ $@ ]========)
	docker build -t appsflyer/terra-crust:$$(git rev-parse --short HEAD) .