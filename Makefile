APP := hive
VERSION := 0.0.1

.PHONY: start build build-ui build-server dev ui-dev server-dev update-env

start: update-env
	cd app && ./server/$(APP)

build: build-ui build-server

build-ui:
	cd ui && pnpm run build

build-server:
	go build -o ./bin/v$(VERSION)/$(APP) ./cmd/main.go
	go build -o ./app/server/$(APP) ./cmd/main.go

dev: 
	make ui-dev & make server-dev

ui-dev:
	cd ui && pnpm run dev

server-dev:
	go run ./cmd/main.go

update-env:
	cp .env ./app/.env

