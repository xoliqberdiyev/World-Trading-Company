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
	@mkdir uploads/partners && mkdir uploads/partners/images/
	@mkdir uploads/categories && mkdir uploads/categories/images && mkdir uploads/categories/icons
	@mkdir uploads/banners && mkdir uploads/banners/images
	@mkdir uploads/news && mkdir uploads/news/images/
	@mkdir uploads/why_us && mkdir uploads/why_us/images/
	@mkdir uploads/certificate && mkdir uploads/certificate/images/
	@mkdir uploads/about_us && mkdir uploads/about_us/images/
	@mkdir uploads/product/ && mkdir uploads/product/images && mkdir uploads/product/banners
	@mkdir uploads/product_medias && mkdir uploads/product_medias/images
	@mkdir uploads/product_file && mkdir uploads/product_file/files