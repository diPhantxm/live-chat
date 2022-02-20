.PHONY: build

build:
	go build -v -o ./.bin/app/ ./cmd/app/

run:
	make build
	go run -v ./cmd/app/

DEFAULT_GOAL: run