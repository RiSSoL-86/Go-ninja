-include src/.env
export DATABASE_URL
export POSTGRES_DB
export POSTGRES_USER
export POSTGRES_PASSWORD

install-pre-commit:
	pre-commit install

ATLAS_BIN = docker run --rm \
	-v "$(CURDIR):/app" \
	-w /app \
	-e DATABASE_URL \
	-e POSTGRES_DB \
	arigaio/atlas

database:
	docker volume create db_go_data
	docker compose -f compose.deps.dev.yml up -d

create-atlas-db:
	@docker exec postgres-go sh -c "createdb -U $(POSTGRES_USER) atlas_dev 2>/dev/null" || echo "Database atlas_dev already exists"

makemigrations:
	@go run src/core/db/atlas_loader/main.go > schema.sql
	@$(ATLAS_BIN) migrate diff $(name) --env local
	@if exist schema.sql del /f /q schema.sql

migrate-up:
	@$(ATLAS_BIN) migrate apply --env local --allow-dirty

migrate-down:
	@$(ATLAS_BIN) migrate down --env local

migrate-status:
	@$(ATLAS_BIN) migrate status --env local

check:
	go mod tidy
	go fmt ./...
	go vet ./...
	go test -race -v ./...

build:
	go build -race -o bin/app.exe ./src

run:
	go run -race ./src
