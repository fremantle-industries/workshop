default: setup.deps test build.compose up

setup.deps:
	go mod tidy

test:
	go test -v -shuffle=on ./...

build.compose:
	docker compose build

build.docker:
	docker build . -t workshop

build.bin:
	go build -ldflags="-w -s" -o ./bin/workshop ./main.go

start.bin:
	./bin/workshop

up:
	docker compose up

down:
	docker compose down

mb: mb.ingest.marketdata mb.catalogs

mb.ingest.marketdata:
	mc mb --ignore-existing local/ingest/marketdata

mb.catalogs:
	mc mb --ignore-existing local/catalogs
