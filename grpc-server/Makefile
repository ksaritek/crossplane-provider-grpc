LOCAL_BIN := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))/.bin

GOIMPORTS_VERSION=v0.8.0
SQLBOILER_VERSION=v4.14.2
MIGRATE_VERSION=v4.15.2
URLENC_VERSION=v1.1.1
GRPCURL_VERSION=v1.8.7

export PATH := ${LOCAL_BIN}:$(PATH)

.PHONY: grpcurl-install
grpcurl-install:
	GOBIN=$(LOCAL_BIN) go install github.com/fullstorydev/grpcurl/cmd/grpcurl@$(GRPCURL_VERSION)

