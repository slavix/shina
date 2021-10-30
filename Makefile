
composeDevFile = docker-compose-dev.yml

init:
	@cp .env-example .env

-include .env

build-dev:
	@docker-compose -f $(composeDevFile) build

build-dev-no-cache:
	@docker-compose  -f $(composeDevFile) build --no-cache

up:
	@docker-compose  -f $(composeDevFile) up -d

down:
	@docker-compose  -f $(composeDevFile) down

enter:
	@docker exec -it --user www-data $(APP_NAME) bash

run:
	@go run cmd/app/main.go