SHELL := /bin/bash

.SILENT: shell
.PHONY: init

init:
	go mod download
	go install github.com/swaggo/swag/cmd/swag@v1.7.3

format:
	gofmt -w .

gendoc:
	${HOME}/go/bin/swag init

run: gendoc
	go run main.go

build: gendoc
	go build .

test:
	go test -cover tests/*.go

ci-bundle: format test

shell:
	echo "Loading .env environment variables..."
	set -a
	source .env
	set +a
