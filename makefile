.PHONY: build

build:
	go build -v -o ./.bin/hospital-rest-api/ ./cmd/hospital-rest-api/

run:
	make build
	go run -v ./cmd/hospital-rest-api/

DEFAULT_GOAL: run