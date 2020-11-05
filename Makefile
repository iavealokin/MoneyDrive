.PHONY: build
build:
	go build -v ./cmd/apiserver

.DEFAULT := build