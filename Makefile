include infrastructure/.env

.PHONY:up
up:
	docker compose -f ./infrastructure/docker-compose.yaml up -d

.PHONY:down
down:
	docker compose -f ./infrastructure/docker-compose.yaml down -v --remove-orphans

.PHONY: build
build:
	go build -o tmp/main cmd/main.go

.PHONY: test
test:
	go test ./internal/...

.PHONY: prepare
prepare:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/vektra/mockery/v2@latest

.PHONY: mockery
mockery:
	mockery

.PHONY: migrate_create
migrate_create:
	migrate create -ext sql -dir database/migration/ -seq migration

.PHONY: migrate_up
migrate_up:
	migrate -path database/migration/ -database "mysql://${DB_USER_NAME}:${DB_USER_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?" -verbose up

migrate_down:
	migrate -path database/migration/ -database "mysql://${DB_USER_NAME}:${DB_USER_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?" -verbose down

