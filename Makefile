include .env
export

build:
	go build -o dist/server main.go
	cp -r templates dist/
	cp -r storages dist/
	cp -r $$CASBIN_MODEL dist/
	cp .env dist/

run: build
	./server

dev:
	go run main.go

watch:
	reflex -s -r '\.go$$' make dev
