#!/usr/bin/env bash

dir_name="proto"

for entry in "$dir_name"/*; do
    echo ${entry}
    echo `protoc -I. ${entry} --go_out=plugins=grpc:$GOPATH/src`
done

echo `go run contracts/main.go`