SHELL := /bin/bash

# Project-wide Makefile for common Go tasks
# Usage: make <target>

BINARY_NAME := go-detail-product-api
PACKAGE := ./...

.PHONY: help run run-dev test tidy deps

help:
	@echo "Makefile commands:"
	@echo "  make run        -> Run the app (go run main.go)"
	@echo "  make run-dev    -> Run the app with race detector"
	@echo "  make test       -> Run all tests"
	@echo "  make tidy       -> Run go mod tidy"
	@echo "  make deps       -> Download modules"
	@echo "  make lint       -> Run golangci-lint (if installed)"

run:
	@echo "Running app..."
	@go run main.go

run-dev:
	@echo "Running app (race detector)..."
	@go run -race main.go

test:
	@echo "Running tests..."
	@go test $(PACKAGE)

tidy:
	@echo "Running go mod tidy..."
	@go mod tidy

deps:
	@echo "Downloading modules..."
	@go mod download
