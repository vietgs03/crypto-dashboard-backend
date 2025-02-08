DB_SOURCE := $(DB_SOURCEE)

dev:
	go run ./cmd/market-data-service
sql:
	sqlc generate

# example run: make new_migration name=00001_init_schema
new_migration:
	migrate create -ext sql -dir sql/schema -seq $(name)

migrate_up:
	migrate -path sql/schema -database "postgresql://postgres:Kiloma123@@localhost:5432/Crypto?sslmode=disable" -verbose up

migrate_down:
	migrate -path sql/schema -database DB_SOURCE="$(DB_SOURCEE)" -verbose down


.PHONY: dev sql new_migration migrate_up migrate_down
