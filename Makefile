.PHONY: all airbloc generate-proto clean install uninstall run test test-all
DEST = $(shell pwd)/build/bin

PROTO_DIR := proto
PROTO_SRCS := $(shell find $(PROTO_DIR) -name *.proto)
RPC_PROTO_SRCS := $(shell find $(PROTO_DIR)/rpc -name *.proto)

all: airbloc bootnode

airbloc:
	./env.sh go install ./cmd/airbloc
	@echo "$(DEST)/airbloc"

bootnode:
	./env.sh go install ./cmd/bootnode
	@echo "$(DEST)/bootnode"

clean:
	@rm -rf build/

install: airbloc
	@cp -f $(DEST)/airbloc $GOPATH/bin/

generate-bind:
	@go run contracts/generate_adapter.go

generate-proto:
	@for PROTO in $(PROTO_SRCS); \
		do protoc -I/usr/local/include -I. \
			-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--go_out=plugins=grpc:$$GOPATH/src \
			--grpc-gateway_out=logtostderr=true:. \
			$$PROTO; \
	done

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
