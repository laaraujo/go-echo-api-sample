api_container = api
local_db_url = postgresql://postgres:postgres@localhost:5432/postgres
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
	@echo goose/rollbal   : rollback latest applied migration on local container
	@echo sqlc/generate   : generate go queries/models with sqlc 

setup:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

build: down
	$(compose_cmd) build --no-cache

down:
	$(compose_cmd) down

run:
	$(compose_cmd) up

stop:
	$(compose_cmd) stop

goose/create:
	goose -dir $(migrations_dir) create rename_this_file sql

goose/status:
	goose -dir $(migrations_dir) postgres "$(local_db_url)" status

goose/migrate:
	goose -dir $(migrations_dir) postgres "$(local_db_url)" up

goose/rollback:
	goose -dir $(migrations_dir) postgres "$(local_db_url)" down

sqlc/generate:
	sqlc generate
