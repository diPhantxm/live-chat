.PHONY: build

build:
	go build -v -o ./.bin/app/ ./cmd/app/

run:
	make build
	./.bin/app/app

DEFAULT_GOAL: run