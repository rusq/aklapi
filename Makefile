SHELL=/bin/sh

IMAGE=aklapi

SRC=main.go $(wildcard aklapi/*.go)

server: $(SRC) 
	go build -o $@

docker:
	docker build -t $(IMAGE) .
.PHONY:docker
