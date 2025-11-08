SHELL := /bin/bash

# Project-wide Makefile for common Go tasks
# Usage: make <target>

BINARY_NAME := go-detail-product-api
PACKAGE := ./...

.PHONY: help run run-dev test tidy deps

help:
	@echo "Makefile commands:"
	@echo "  make run        -> Run the app (go run main.go)"
	@echo "  make test       -> Run all tests"
	@echo "  make tidy       -> Run go mod tidy"
	@echo "  make deps       -> Download modules"
	@echo "  make lint       -> Run golangci-lint (if installed)"

run:
	@echo "Running app..."
	@go run main.go

test:
	@echo "Running tests..."
	go test -v -cover -short ./...

tidy:
	@echo "Running go mod tidy..."
	@go mod tidy

deps:
	@echo "Downloading modules..."
	@go mod download
