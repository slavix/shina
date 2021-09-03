
composeDevFile = docker-compose-dev.yml

build-web-dev:
	@docker-compose -f $(composeDevFile) build

build-web-dev-no-cache:
	@docker-compose  -f $(composeDevFile) build --no-cache

up-web-dev:
	@docker-compose  -f $(composeDevFile) up -d

down-web-dev:
	@docker-compose  -f $(composeDevFile) down

run:
	@go run cmd/app/main.go