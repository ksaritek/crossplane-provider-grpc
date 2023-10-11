LOCAL_BIN := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))/.bin

GOIMPORTS_VERSION=v0.8.0
SQLBOILER_VERSION=v4.14.2
MIGRATE_VERSION=v4.15.2
URLENC_VERSION=v1.1.1
GRPCURL_VERSION=v1.8.7

export PATH := ${LOCAL_BIN}:$(PATH)

### General
.PHONY: all
all: clean tools build test

.PHONY: tools
tools: go-imports tidy grpcurl-install

### Golang
.PHONY: go-imports
go-imports:
	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION)

.PHONY: grpcurl-install
grpcurl-install:
	GOBIN=$(LOCAL_BIN) go install github.com/fullstorydev/grpcurl/cmd/grpcurl@$(GRPCURL_VERSION)

.PHONY: build
build:
	go build -o $(LOCAL_BIN)/$(BINARY_NAME) -a -race ./cmd/$(BINARY_NAME)

.PHONY: build-fast
build-fast:
	go build -o $(LOCAL_BIN)/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

.PHONY: run
run:
	go run main.go

.PHONY: test
test:
	go test -failfast -v -vet=off -race ./...

.PHONY: update
update: tidy
	go get -u ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: imports
imports:
	$(LOCAL_BIN)/goimports -l -w .

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: clean
clean:
	rm -rf $(LOCAL_BIN)