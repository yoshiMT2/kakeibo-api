ifneq (,$(wildcard ./.env))
    include .env
    export
endif

build:
	@go build -o bin/kakeibo-api

run: build
	@./bin/kakeibo-api

test:
	@go test -v ./...

migrate:
	migrate -path migrations/ -database "${DATABASE_URL}" up
