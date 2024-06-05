include .env

install:
	@$(MAKE) -s down
#	@$(MAKE) -s docker-build
	@$(MAKE) -s up
	@$(MAKE) -s migrate-up
	@echo "--- Application installed ---"

up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

docker-build:
	docker build -t app-web .

shell:
	docker exec -it ${HTTP_SERVER_NAME} /bin/sh -c '$(cmd)'

migrate-create:
	migrate create -ext sql -dir migrations $(name)

migrate-up:
	migrate -source $(MIGRATION_URL) -database $(DB_DRIVER_NAME)://$(DB_USER_NAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE) -verbose up

migrate-down:
	migrate -source $(MIGRATION_URL) -database $(DB_DRIVER_NAME)://$(DB_USER_NAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE) -verbose down