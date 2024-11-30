.PHONY: build run test

build:
	mkdir -p bin
	go build -o bin/app.exe

run: build
	./bin/app.exe

test:
	go test -v ./... -count=1