DB_DSN = "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
MIGRATE = migrate -path ./migrations -database $(DB_DSN)

.PHONY: help migrate-new migrate-up migrate-down

help:
	@echo Available commands:
	@echo   make migrate-new NAME=name  - Create new migration
	@echo   make migrate-up            - Apply all migrations
	@echo   make migrate-down          - Rollback last migration

migrate-new:
	@if not defined NAME ( \
		echo NameMigrate: make migrate-new NAME=NameMigrate && exit 1 \
	)
	migrate create -ext sql -dir ./migrations -seq $(NAME)

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE)
run:
	go run cmd/main.go
