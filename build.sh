#!/usr/bin/env bash

dests=(
    "types"
    "data"
    "exchange"
    "schemas"
    "collections"
    "producer"
    "warehouse"
)

for dest in "${dests[@]}"; do
    echo `protoc -I proto proto/${dest}.proto --go_out=plugins=grpc:$GOPATH/src`
done

echo `go run contracts/main.go`