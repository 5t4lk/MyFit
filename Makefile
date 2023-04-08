.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go
run: build
	./.bin/bot

build-image:
	docker build -t myfit-bot-diplom:v0.1 .

start-container:
	docker run --name myfit-bot --env-file .env myfit-bot-diplom:v0.1
