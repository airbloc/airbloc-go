#!/usr/bin/env bash

protoc -I proto/ proto/producer.proto --go_out=plugins=grpc:proto