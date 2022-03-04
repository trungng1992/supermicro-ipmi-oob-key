GOCMD ?= go
GOBUILD = $(GOCMD) build
BINARY_NAME = supermicro-ipmi-oob-key

GO_LINUX := GO111MODULE=on GOOS=linux GOARCH=amd64 go
GO_DARWIN := GO111MODULE=on GOOS=darwin GOARCH=amd64 go
GO_WINDOWS := GO111MODULE=on GOOS=windows GOARCH=amd64 go


ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

default: build
build:
	$(GO_LINUX) build -o bin/$(BINARY_NAME)-linux
	$(GO_DARWIN) build -o bin/$(BINARY_NAME)-darwin
	$(GO_WINDOWS) build -o bin/$(BINARY_NAME)-windows