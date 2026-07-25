MIGRATIONS_PATH = ./cmd/migrate/migrations
# TODO заменить на подтягивание из конфига
DATABASE_CONNSTR = "postgres://postgres:postgres@localhost/demos?sslmode=disable"


.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DATABASE_CONNSTR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DATABASE_CONNSTR) down $(filter-out $@,$(MAKECMDGOALS))