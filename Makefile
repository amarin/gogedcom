# Go parameters
GOCMD=GO111MODULE=on GOFLAGS=-mod=vendor go
DEPCMD=$(GOCMD) mod

# Detect version by current commit tag
ifeq ($(OS), Windows_NT)
	GIT_VERSION := $(shell git describe --tags --first-parent --abbrev=0 2>nil)
else
	GIT_VERSION := $(shell git describe --tags --first-parent --abbrev=0 2>/dev/null)
endif

GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

.PHONY: all deps build clean

.PHONY: test
test: ## test project
	export GOPATH=$GOPATH:`pwd`
	$(GOTEST) -v ./...

.PHONY: tidy
tidy: ## Fix dependencies records in go.mod
	@go mod tidy

deps: tidy ## fix dependencies
	$(DEPCMD) vendor

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'