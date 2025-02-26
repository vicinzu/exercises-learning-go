.DEFAULT_GOAL:= build

.PHONY:fmt vet build

clean:
	rm -rf ./bin/*

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o bin/ ./...

run: build
	./bin/$(shell basename $(CURDIR))