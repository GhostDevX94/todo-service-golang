# Load environment variables from .env file
# If .env doesn't exist, copy from env.example
ifneq (,$(wildcard .env))
    include .env
else
    ifneq (,$(wildcard env.example))
        $(shell cp env.example .env)
        include .env
    endif
endif

# Set default values only for variables that are not in .env
MIGRATE ?= /opt/homebrew/bin/migrate


build:
	@echo "Putting the application together..."
	go build -o $(BINARY_PATH) cmd/main.go

clean:
	@echo "Removing binary file..."
	rm -f $(BINARY_PATH)

run: build
	@echo "Starting $(BINARY_PATH) on port $(APP_PORT)..."
	@$(BINARY_PATH) &

stop:
	@echo "Stopping $(BINARY_NAME)..."
	@pkill -f "$(BINARY_PATH)" || true

restart: stop build run
	@echo "Restart completed."

create-migration:
	@read -p "Enter a migration name: " NAME; \
	if [ -z "$$NAME" ]; then echo "Migration name must not be empty!"; exit 1; fi; \
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) -seq "$$NAME"

migrate-up:
	@echo "Applying new migrations..."
	$(MIGRATE) -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) up

migrate-down:
	@echo "Rolling back the last migration..."
	$(MIGRATE) -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) down 1

migrate-steps:
	@read -p "Enter the number of steps:" STEPS; \
	if [ -z "$$STEPS" ]; then echo "Steps not specified!"; exit 1; fi; \
	$(MIGRATE) -database $(DATABASE_URLL) -path $(MIGRATIONS_DIR) up $$STEPS

migrate-rollback:
	@read -p "Enter the number of steps to roll back: " STEPS; \
	if [ -z "$$STEPS" ]; then echo "Steps not specified!"; exit 1; fi; \
	$(MIGRATE) -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) down $$STEPS


