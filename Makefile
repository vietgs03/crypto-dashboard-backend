dev:
	go run ./cmd/market-data-service
sql:
	sqlc generate

.PHONY: dev sql
