api_container = api
migrations_dir = db/migrations
compose_cmd = docker compose -f docker/compose.yml

help:
	@echo build           : build local docker setup
	@echo down            : remove local containers
	@echo run             : spin local containers up
	@echo stop            : stop local containers
	@echo goose/create    : create a new migration file (with a default name, to be renamed)
	@echo goose/validate  : validate migration files
	@echo goose/migrate   : run migrations on local container
	@echo sqlc/generate   : generate go queries/models with sqlc 

build: down
	$(compose_cmd) build --no-cache

down:
	$(compose_cmd) down

run:
	$(compose_cmd) up

stop:
	$(compose_cmd) stop

goose/create:
	$(compose_cmd) run $(api_container) goose -dir $(migrations_dir) create rename_this_file sql

goose/validate:
	$(compose_cmd) run $(api_container) goose -dir $(migrations_dir) validate

goose/migrate:
	$(compose_cmd) run $(api_container) goose -dir $(migrations_dir) up 

sqlc/generate:
	$(compose_cmd) run $(api_container) sqlc generate
