GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && b=`basename $$PWD` && echo $$b/$$a)
APP_NAME=$(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")

.PHONY: init
# init env
init:
	go install github.com/google/wire/cmd/wire@latest

.PHONY: generate
# generate
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: docs
docs:
	swag init --pd -g ./cmd/main.go --output ./docs --parseDepth 3

.PHONY: serv
serv:
	cd cmd && go run .

.PHONY: wire
wire:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: data
data:
	go mod tidy
	cd ./internal/data && go generate ./ent

.PHONY: asynqmon
asynqmon:
	docker run --rm --name asynqmon --network dev_dev_net -p 3000:3000 hibiken/asynqmon --port=3000 --redis-addr=redis-dev:6379 --redis-password=WGlhb2h1b3poaTIwMjIu