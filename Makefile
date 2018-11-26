.PHONY: build test

build:
	go build -o dgd cmd/dgd/main.go

test:
	go test

