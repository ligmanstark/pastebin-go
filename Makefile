ifneq (,$(wildcard .env))
    include .env
    export
endif

deploy-prod:
	docker compose -f docker-compose.prod.yml up -d

down-prod:
	docker compose -f docker-compose.prod.yml down

deploy-local:
	docker-compose -f docker-compose.local.yml up -d

down-local:
	docker-compose -f docker-compose.local.yml down

migrate-up:
	docker run --rm --network pastebin_pastebin_net -v $(PWD)/migrations:/migrations --env-file .env migrate/migrate -path /migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@db:5432/$(DB_DBNAME)?sslmode=$(DB_SSLMODE)" up

migrate-down:
	docker run --rm --network pastebin_pastebin_net -v $(PWD)/migrations:/migrations --env-file .env migrate/migrate -path /migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@db:5432/$(DB_DBNAME)?sslmode=$(DB_SSLMODE)" down 1

migrate-force:
	docker run --rm --network pastebin_pastebin_net -v $(PWD)/migrations:/migrations --env-file .env migrate/migrate -path /migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@db:5432/$(DB_DBNAME)?sslmode=$(DB_SSLMODE)" force $(version)

migrate-version:
	docker run --rm --network pastebin_pastebin_net -v $(PWD)/migrations:/migrations --env-file .env migrate/migrate -path /migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@db:5432/$(DB_DBNAME)?sslmode=$(DB_SSLMODE)" version

migrate-create:
	docker run --rm -v $(PWD)/migrations:/migrations migrate/migrate create -ext sql -dir /migrations -seq $(name)

connect-db:
	docker run -it --rm --network pastebin_pastebin_net --env-file .env postgres psql -h db -U $(DB_USER) -d $(DB_DBNAME)

create-dump:
	docker exec -t pastebin-db-1 pg_dump -U $(DB_USER) $(DB_DBNAME) > dumps/dump.sql