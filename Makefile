PKG = cmd

default:
	build

build:
	go build $(PKG)/main.go

test:
	go test
	go test ./impl

.PHONY: build