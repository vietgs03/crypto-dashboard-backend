# dev:
# 	go run ./cmd/market-data-service
# sql:
# 	sqlc generate

# # example run: make new_migration name=00001_init_schema
# new_migration:
# 	migrate create -ext sql -dir sql/schema -seq $(name)

# migrate_up:
# 	migrate -path sql/schema -database "postgresql://postgres:Kiloma123@@localhost:5432/Crypto?sslmode=disable" -verbose up

# migrate_down:
# 	migrate -path sql/schema -database "postgresql://postgres:Kiloma123@@localhost:5432/Crypto?sslmode=disable" -verbose down


WORKER_IMAGE=1.24.0-alpine3.21

DIRS := internal/

DONT_STOP := db redis

.PHONY: tsl-generate build-services stop-services start-services reset-services doc-generate swagger-2-to-3


build-specific-services:
	@echo "Building services: $(filter-out $@,$(MAKECMDGOALS))"
	@for service in $(filter-out $@,$(MAKECMDGOALS)); do \
		cp -R ./shared ./$$service; \
		rm -rf ./$$service/shared/tests; \
	done
	docker-compose -p easier-be build $(filter-out $@,$(MAKECMDGOALS))
	@for service in $(filter-out $@,$(MAKECMDGOALS)); do \
		rm -rf ./$$service/shared; \
	done


start-specific-services:
	@echo "Starting services: $(filter-out $@,$(MAKECMDGOALS))"
	@for service in $(filter-out $@,$(MAKECMDGOALS)); do \
		cp -R ./shared ./$$service; \
		rm -rf ./$$service/shared/tests; \
	done
	docker-compose up -d $(filter-out $@,$(MAKECMDGOALS))
	@for service in $(filter-out $@,$(MAKECMDGOALS)); do \
		rm -rf ./$$service/shared; \
	done
	docker-compose logs $(filter-out $@,$(MAKECMDGOALS)) -f


reset-specific-services:
	make stop-services
	make start-specific-services $(filter-out $@,$(MAKECMDGOALS))


stop-services:
	docker-compose down


go-lint-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	cp hooks/pre-commit .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit


go-lint:
	@for dir in $(DIRS); do \
		echo "Running golangci-lint in $$dir"; \
		cd $$dir && golangci-lint run && cd ..; \
	done