PKG = cmd

default:
	build-main

build-main:
	go build $(PKG)/main.go

.PHONY: build-main