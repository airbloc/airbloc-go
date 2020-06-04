export GO111MODULE=on

.PHONY: deps test test-all test-e2e
DEST = $(shell pwd)/build/bin

# test runner (can be overriden by CI)
GOTEST ?= go test

# build tags
LDFLAGS += -X "main.Version=$(shell git rev-list --tags --max-count=1)"
LDFLAGS += -X "main.BuildDate=$(shell date)"
LDFLAGS += -X "main.GitCommit=$(shell git rev-parse --short HEAD)"
LDFLAGS += -X "main.GitBranch=$(shell git symbolic-ref -q --short HEAD)"

deps:
ifeq ($(shell which mockgen), )
	@echo "Installing Dependency: mockgen"
	@go install github.com/golang/mock/mockgen
endif

test: test-all

test-all:
	@$(GOTEST) -v -count 1 `go list ./... | grep -v test/e2e`

test-e2e:
	@$(GOTEST) -v -count 1 `go list ./test/e2e` $(FLAGS)
