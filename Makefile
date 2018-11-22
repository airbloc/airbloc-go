.PHONY: all airbloc generate-proto clean install uninstall run test test-all
DEST = $(shell pwd)/build/bin

PROTO_DIR := proto
PROTO_SRCS := $(wildcard $(PROTO_DIR)/*.proto)

all: airbloc

airbloc:
	./env.sh go install ./cmd/airbloc
	@echo "$(DEST)/airbloc"

clean:
	@rm -rf build/

install: airbloc
	@cp -f $(DEST)/airbloc $GOPATH/bin/

generate-bind:
	@go run contracts/generate_adapter.go

generate-proto:
	@for PROTO in $(PROTO_SRCS); do protoc -I. $$PROTO --go_out=plugins=grpc:$$GOPATH/src; done
    @protoc --go_out=. proto/p2p/message.proto

uninstall:
	@rm -f $GOPATH/bin/airbloc

run: airbloc
	@$(DEST)/airbloc

test: test-all

test-all:
	@go test -v ./...
