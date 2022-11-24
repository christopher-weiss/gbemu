.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

test: 
	go test ./...
.PHONY:test

build: vet
	go build -o gbemu ./...
.PHONY:build
