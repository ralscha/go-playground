SHELL:=/bin/bash

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## upgrade-libraries: upgrade all libraries and run tidy and verify
.PHONY: upgrade-libraries
upgrade-libraries:
	go get go@latest
	go get -u ./...
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go mod verify

## build: build s3copy tool
build:
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o s3copy .

