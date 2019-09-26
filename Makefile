export GO111MODULE=on

.PHONY: all deps airbloc generate-proto clean install uninstall run test test-all contracts
DEST = $(shell pwd)/build/bin

PROTO_DIR := proto
PROTO_SRCS := $(shell find $(PROTO_DIR) -name *.proto)
RPC_PROTO_SRCS := $(shell find $(PROTO_DIR)/rpc -name *.proto)

# test runner (can be overriden by CI)
GOTEST ?= go test

# build tags
LDFLAGS += -X "main.Version=$(shell git rev-list --tags --max-count=1)"
LDFLAGS += -X "main.BuildDate=$(shell date)"
LDFLAGS += -X "main.GitCommit=$(shell git rev-parse --short HEAD)"
LDFLAGS += -X "main.GitBranch=$(shell git symbolic-ref -q --short HEAD)"

all: airbloc bootnode

deps:
ifeq ($(shell which mockgen), )
	@echo "Installing Dependency: mockgen"
	@go install github.com/golang/mock/mockgen
endif

airbloc:
	@go build -o "$(DEST)/airbloc" -ldflags '${LDFLAGS}' ./cmd/airbloc
	@echo "$(DEST)/airbloc"

bootnode:
	@go build -o "$(DEST)/bootnode" -ldflags '${LDFLAGS}' ./cmd/bootnode
	@echo "$(DEST)/bootnode"

clean:
	@rm -rf build/

install: all
	@cp -f $(DEST)/* $$GOPATH/bin/

contracts:
	@cd contracts; npm run compile

generate-bind: contracts
	@go run contracts/generate_adapter.go

generate-proto:
	@for PROTO in $(PROTO_SRCS); do \
	  protoc -I/usr/local/include -I. \
			--go_out=plugins=grpc:$$GOPATH/src \
			$$PROTO; \
	done

docs:
	@mkdir -p build/docs
	@for VERSION in $(PROTO_DIR)/rpc/*; do \
		for VARIANTS_PATH in $$VERSION/*; do \
			VARIANT=`echo $$VARIANTS_PATH | rev | cut -d/ -f1 | rev`; \
			protoc -I/usr/local/include -I. \
				-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
				--doc_out=./build/docs/ \
				--doc_opt=markdown,$$VARIANT.md \
				$$VARIANTS_PATH/*.proto; \
		done; \
	done;

generate-python-pb:
	@mkdir -p build/gen
	@python -m grpc_tools.protoc -I. \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--python_out=./build/gen \
		--grpc_python_out=./build/gen \
		$(RPC_PROTO_SRCS)

uninstall:
	@rm -f $GOPATH/bin/airbloc

run: airbloc
	@$(DEST)/airbloc

test: test-all

test-all:
	@$(GOTEST) -v -count 1 `go list ./... | grep -v test/e2e`

test-e2e:
	@$(GOTEST) -v -count 1 `go list ./test/e2e` $(FLAGS)
