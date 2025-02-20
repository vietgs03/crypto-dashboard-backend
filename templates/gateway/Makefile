.PHONY: migrate-add migrate-run migrate-down di-generate test test-silent test-cov build build-init go-lint

migrate-add:
	@read -p "Enter migration name: (example: create-todo-table) " name; \
		go run ./cmd/migrate create $$name sql 

migrate-run:
	go run ./cmd/migrate up

migrate-down:
	go run ./cmd/migrate down

di-generate:
	wire ./app

go-lint:
	golangci-lint run

# Test
test:
	go test -v ./...
test-silent:
	go test ./...

test-cov:
	go test -cover ./...

build:
	cp -R ../shared .
	rm -rf ./shared/tests
	docker build -t easier-naas -f Dockerfile.prd .
	rm -R shared/

build-init:
	cp -R ../shared .
	rm -rf ./shared/tests
	docker buildx build --platform linux/amd64 --load -t easier-naas:latest -f Dockerfile.prd .
	rm -R shared/

generate-proto:
	protoc --go_out=./proto --go-grpc_out=./proto proto/schema/*.proto