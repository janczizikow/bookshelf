include .env
export

.PHONY: run
run:
	go run ./cmd/api

.PHONY: psql
psql:
	psql $(DB_DSN)