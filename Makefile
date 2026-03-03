SHELL=/bin/sh

IMAGE=aklapi

SRC=$(wildcard aklapi/*.go)
PKG=./cmd/aklapi

server: $(SRC) 
	go build -o $@ $(PKG)

test:
	go test ./... -race
.PHONY: test

docker:
	docker build -t $(IMAGE) .
.PHONY:docker
