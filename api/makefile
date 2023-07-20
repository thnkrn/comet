SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage
DOCKERCMD=docker
DOCKER_IMAGE_NAME=comet-api
CONTAINER_REGISTRY_HOST=asia.gcr.io
CONTAINER_REGISTRY_PROJECT_ID=tdg-analytics-poc
CONTAINER_REGISTRY_IMAGE_NAME=comet-api

all: test build

${BINARY_DIR}:
	mkdir -p $(BINARY_DIR)

build: ${BINARY_DIR} ## Compile the code, build Executable File
	$(GOCMD) build -o $(BINARY_DIR) -v ./cmd/api

run: ## Start application
	$(GOCMD) run ./cmd/api

test: ## Run tests
	$(GOCMD) test ./... -cover -v

test-coverage: ## Run tests and generate coverage file
	$(GOCMD) test ./... -coverprofile=$(CODE_COVERAGE).out
	$(GOCMD) tool cover -html=$(CODE_COVERAGE).out

deps: ## Install dependencies
	$(GOCMD) install github.com/google/wire/cmd/wire@latest
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	$(GOCMD) get -u -t -d -v ./...
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

deps-cleancache: ## Clear cache in Go module
	$(GOCMD) clean -modcache

wire: ## Generate wire_gen.go
	cd pkg/di && wire

docker-build: ## Build docker image with default setting and platform
	$(DOCKERCMD) build -t $(CONTAINER_REGISTRY_HOST)/$(CONTAINER_REGISTRY_PROJECT_ID)/$(CONTAINER_REGISTRY_IMAGE_NAME):$(tag) .

docker-build-amd: ## Build docker image with default setting and platform as amd64
	$(DOCKERCMD) buildx build --platform=linux/amd64 -t $(CONTAINER_REGISTRY_HOST)/$(CONTAINER_REGISTRY_PROJECT_ID)/$(CONTAINER_REGISTRY_IMAGE_NAME):$(tag) .

docker-run: ## Run docker image locally
	$(DOCKERCMD) run --rm -it -v /tmp/rocks-dev:/tmp/rocks-dev -p 8080:8080 $(CONTAINER_REGISTRY_HOST)/$(CONTAINER_REGISTRY_PROJECT_ID)/$(CONTAINER_REGISTRY_IMAGE_NAME):$(tag)

docker-push: ## Push the image to Container Registry (EX. tag=0.0.1)
	$(DOCKERCMD) push $(CONTAINER_REGISTRY_HOST)/$(CONTAINER_REGISTRY_PROJECT_ID)/$(CONTAINER_REGISTRY_IMAGE_NAME):$(tag)

docker-deploy: ## Build image by pulling cache from latest build on cloud registry and push to cloud registry
	$(DOCKERCMD) buildx build --platform linux/amd64 -t $(CONTAINER_REGISTRY_HOST)/$(CONTAINER_REGISTRY_PROJECT_ID)/$(CONTAINER_REGISTRY_IMAGE_NAME):$(tag)-amd64  --cache-from type=registry,ref=asia-docker.pkg.dev/gcp-analytics-platform-mgt/product-app/comet-api:buildcache .
	$(DOCKERCMD) push $(CONTAINER_REGISTRY_HOST)/$(CONTAINER_REGISTRY_PROJECT_ID)/$(CONTAINER_REGISTRY_IMAGE_NAME):$(tag)-amd64 

mockery:
	cd pkg/usecase && mockery --all --output=../mocks/usecase --case underscore
	cd ../../
	cd pkg/repository && mockery --all --output=../mocks/repository --case underscore
	cd ../../
	cd pkg/driver/log && mockery --all --output=../../mocks/driver/log --case underscore

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'