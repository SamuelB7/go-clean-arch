include .env
MIGRATIONS_PATH = ./migrate/migrations

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path $(MIGRATIONS_PATH) -database $(DATABASE_URL) up

.PHONY: migrate-down
migrate-down:
	@migrate -path $(MIGRATIONS_PATH) -database $(DATABASE_URL) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: gen-docs
gen-docs:
	@echo "Generating API documentation..."
	@swag init -g ./cmd/api/main.go -o ./docs --parseDependency --parseInternal
	@echo "API documentation generated successfully."