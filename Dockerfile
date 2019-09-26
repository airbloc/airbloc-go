FROM golang:1.13-alpine AS base

# Install build toolchain for alpine
RUN apk add --no-cache make git g++ musl-dev linux-headers bash ca-certificates
WORKDIR /airbloc

# 1. Fetch and cache go module dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# 2. Copy rest of the sources and build it
FROM base AS builder
COPY . .
RUN GOOS=linux GOARCH=amd64 make all

# 3. Pull binary into a clean alpine container
FROM alpine:latest

RUN apk add ca-certificates

COPY --from=builder /airbloc/build/bin/* /usr/local/bin/

EXPOSE 2471-2474
ENTRYPOINT ["airbloc"]
# or it can be ENTRYPOINT ["bootnode"]
