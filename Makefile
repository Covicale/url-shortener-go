# Change these variables as necessary.
MAIN_PACKAGE_PATH := ./cmd
BINARY_NAME := main

build:
	go build -o ./build/$(BINARY_NAME) $(MAIN_PACKAGE_PATH)/$(BINARY_NAME).go

run:
	go run $(MAIN_PACKAGE_PATH)/$(BINARY_NAME).go

create-db:
	docker build -t go-shorten-db -f db.Dockerfile .

run-db:
	docker run -d --name postgres-db -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=covicale -p 5432:5432 go-shorten-db

randr-db: create-db run-db