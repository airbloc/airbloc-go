FROM golang:1.11.5-alpine as base

# Install build toolchain for alpine
RUN apk add --no-cache make git g++ musl-dev linux-headers bash ca-certificates
WORKDIR /airbloc

# use go modules
ENV GO111MODULE=on

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
COPY --from=builder /airbloc/build/bin/* /usr/local/bin/

EXPOSE 9124
ENTRYPOINT ["airbloc"]
# or it can be ENTRYPOINT ["bootnode"]
