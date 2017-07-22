PKG = cmd/cleaningplan-server

build:
	go build ./$(PKG)

test:
	go test ./...

default:
	build

.PHONY: build