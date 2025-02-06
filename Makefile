-include .env

.SILENT:

DB_URL=postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

build:
	@go build -o bin/main main.go

run: build
	@./bin/main

migration:
	@migrate create -ext sql -dir ./migrations -seq $(name)

migrateup:
	@migrate -path ./migrations -database "$(DB_URL)" -verbose up

migratedown:
	@migrate -path ./migrations -database "$(DB_URL)" -verbose down

tidy:
	@go mod tidy
	@go mod vendor

createsuperuser:
	@go run main.go createsuperuser

createfiles:
	@mkdir uploads/ && mkdir uploads/medias/ && mkdir uploads/medias/files/