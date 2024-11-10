.PHONY: build-local build templ notify-templ-proxy run

-include .env

build-local:
	@go build -o ./bin/main cmd/main/main.go

build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main cmd/main/main.go

templ:
	@TEMPL_EXPERIMENT=rawgo templ generate --watch --proxy=http://localhost:$(APP_PORT) --proxyport=$(TEMPL_PROXY_PORT) --open-browser=false --proxybind="0.0.0.0"

notify-templ-proxy:
	@templ generate --notify-proxy --proxyport=$(TEMPL_PROXY_PORT)

run:
	@make templ & sleep 1
	@air

# run: build
# 	@./bin/main

# templ:
# 	@templ generate --watch -proxy=http://localhost:3003

# tailwind:
# 	npx tailwindcss -i assets/css/input.css -o assets/css/tailwind.css --watch

# build:
# 	npx tailwindcss -i assets/css/input.css -o assets/css/tailwind.css
# 	@templ generate
# 	@go build -o bin/main main.go