include .env
export

export PROJECT_ROOT=$(shell pwd)

run:
	@go run cmd/app/main.go

up:
	docker compose up -d --build

down:
	docker compose down

app:
	docker compose up -d --build app

migrate-create:
	@if [ -z "$(seq)" ]; then \
  		exit 1; \
  	fi; \
	docker compose run --rm postgres-migrate \
		create \
		-ext sql \
		-dir /migrations	\
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
  		exit 1; \
  	fi; \
	docker compose run --rm postgres-migrate \
		--path /migrations \
		--database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable \
		"$(action)"

env-up:
	@docker compose up -d postgres redis

env-down:
	@docker compose down postgres redis

env-cleanup:
	@read -p "Delete all volume files? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down postgres && \
		sudo rm -rf out/pgdata && \
		echo "Env files deleted"; \
	else \
		echo "Delete canceled"; \
	fi

env-port-forward-pg:
	@docker compose up -d port-forwarder-postgres

env-port-forward-redis:
	@docker compose up -d port-forwarder-redis

env-port-close-pg:
	@docker compose down port-forwarder-postgres

env-port-close-redis:
	@docker compose down port-forwarder-redis