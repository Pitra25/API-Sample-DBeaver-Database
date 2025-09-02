# Go Makefile
BINARY_NAME=app
GO=go
GO_FILES=$(wildcard *.go)

.PHONY: all build run test clean

all: test build

build:
    $(GO) build -o $(BINARY_NAME) .

run:
    $(GO) run main.go

clean:
    rm -f $(BINARY_NAME)
    rm -f coverage.txt

# Для проектов с CGO (как SQLite)
build-with-cgo:
    CGO_ENABLED=1 $(GO) build -o $(BINARY_NAME) .