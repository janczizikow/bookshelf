include .env
export

.PHONY: run
run:
	go run ./cmd/api

.PHONY: migrate
migrate:
	migrate -path ./internal/database/migrations -database=$(DB_DSN) up

.PHONY: psql
psql:
	psql $(DB_DSN)