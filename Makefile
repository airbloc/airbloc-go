.PHONY: all aerod generate-proto clean install uninstall run test test-all
DEST = $(shell pwd)/build/bin

PROTO_DIR := proto
PROTO_SRCS := $(wildcard $(PROTO_DIR)/*.proto)

all: aerod

aerod:
	./env.sh go install ./cmd/aerod
	@echo "$(DEST)/aerod"

clean:
	@rm -rf build/

install: aerod
	@cp -f $(DEST)/aerod $GOPATH/bin/

generate-proto:
	@for PROTO in $(PROTO_SRCS); do protoc -I. $$PROTO --go_out=plugins=grpc:$$GOPATH/src; done

uninstall:
	@rm -f $GOPATH/bin/aerod

run: aerod
	@$(DEST)/aerod

test: test-all

test-all:
	@go test -v ./...
