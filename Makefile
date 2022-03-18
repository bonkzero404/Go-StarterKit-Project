build:
	go build -o dist/server main.go
	cp -r templates dist/
	cp -r storages dist/
	cp -r casbin_models dist/
	cp .env dist/

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run
