export GO111MODULE=on

.PHONY: all deps airbloc generate-proto clean install uninstall run test test-all contracts
DEST = $(shell pwd)/build/bin

PROTO_DIR := proto
PROTO_SRCS := $(shell find $(PROTO_DIR) -name *.proto)
RPC_PROTO_SRCS := $(shell find $(PROTO_DIR)/rpc -name *.proto)

all: airbloc bootnode

deps:
	@go build -v ...
	@go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

airbloc:
	@go build -o "$(DEST)/airbloc" ./cmd/airbloc
	@echo "$(DEST)/airbloc"

bootnode:
	@go build -o "$(DEST)/bootnode" ./cmd/bootnode
	@echo "$(DEST)/bootnode"

clean:
	@rm -rf build/

install: airbloc
	@cp -f $(DEST)/airbloc $GOPATH/bin/

contracts:
	@cd contracts; npm run compile

generate-bind: contracts
	@go run contracts/generate_adapter.go

generate-proto:
	@for PROTO in $(PROTO_SRCS); do \
	  protoc -I/usr/local/include -I. \
			-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--go_out=plugins=grpc:$$GOPATH/src \
			--grpc-gateway_out=logtostderr=true:$$GOPATH/src \
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
	@go test -v ./...
