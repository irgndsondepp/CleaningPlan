PKG = cmd

default:
	build-main

build-main:
	go build $(PKG)/main.go

test:
	go test
	go test ./people
	go test ./people/tasks

.PHONY: build-main