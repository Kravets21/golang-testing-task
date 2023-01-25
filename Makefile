.PHONY: all test clean

SHELL = /bin/bash

APP_CONTAINER_NAME := task_api
docker_bin := $(shell command -v docker 2> /dev/null)
docker_compose_bin := $(shell command -v docker-compose 2> /dev/null)

build: ## Build all containers and start (interact) for development
	@echo "build task_api api service..."
	$(docker_compose_bin) build --no-cache --parallel --force-rm
	$(docker_compose_bin) up --remove-orphans --force-recreate

up: ## Start all containers (no interact) for development
	@echo "build and run..."
	$(docker_compose_bin) up --no-recreate -d

down: ## Stop all started for development containers
	@echo "down..."
	$(docker_compose_bin) down

run:
	go run cmd/api/main.go

test-docker:
	$(docker_compose_bin) exec -it $(APP_CONTAINER_NAME) go test ./test -v -race

test:
	go test ./test -v -race