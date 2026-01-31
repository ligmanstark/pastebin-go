ifneq (,$(wildcard .env))
    include .env
    export
endif

deploy-prod:
	docker compose -f docker-compose.prod.yml up -d --build

down-prod:
	docker compose -f docker-compose.prod.yml down -v

deploy-local:
	docker-compose -f docker-compose.local.yml up -d --build

down-local:
	docker-compose -f docker-compose.local.yml down -v

migrate-up:
	docker run --rm --network $(network_name) -v $(PWD)/migrations:/migrations --env-file .env migrate/migrate -path /migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@db:5432/$(DB_DBNAME)?sslmode=$(DB_SSLMODE)" up

migrate-down:
	docker run --rm --network $(network_name) -v $(PWD)/migrations:/migrations --env-file .env migrate/migrate -path /migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@db:5432/$(DB_DBNAME)?sslmode=$(DB_SSLMODE)" down 1

migrate-force:
	docker run --rm --network $(network_name) -v $(PWD)/migrations:/migrations --env-file .env migrate/migrate -path /migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@db:5432/$(DB_DBNAME)?sslmode=$(DB_SSLMODE)" force $(version)

migrate-version:
	docker run --rm --network $(network_name) -v $(PWD)/migrations:/migrations --env-file .env migrate/migrate -path /migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@db:5432/$(DB_DBNAME)?sslmode=$(DB_SSLMODE)" version

migrate-create:
	docker run --rm -v $(PWD)/migrations:/migrations migrate/migrate create -ext sql -dir /migrations -seq $(name)

connect-db:
	docker run -it --rm --network $(network_name) --env-file .env postgres psql -h db -U $(DB_USER) -d $(DB_DBNAME)

create-dump:
	docker exec -t pastebin-db-1 pg_dump -U $(DB_USER) $(DB_DBNAME) > dumps/dump.sql

accept-dump:
	cat dumps/dump.sql | docker exec -i pastebin-db-1 psql -U $(DB_USER) -d $(DB_DBNAME)