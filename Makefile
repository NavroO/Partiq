# -------------------
# CONFIG
# -------------------
include apps/api/.env
export
MIGRATIONS_DIR=apps/api/migrations
SEED=apps/api/seed/seed.go
API=apps/api
CLIENT=apps/client

# -------------------
# DEV
# -------------------

dev:
	npx concurrently -n API,FE -c yellow,cyan \
	"cd apps/api && go run cmd/api/main.go" \
	"cd apps/client && npm run dev"

# -------------------
# SEEDING
# -------------------

seed:
	@echo "🌱 Running seed script..."
	cd $(API) && go run $(SEED)

# -------------------
# MIGRATIONS
# -------------------

migration-add:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_DIR) $$name

migration-run:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

migration-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

migration-force:
	@read -p "Enter version to force: " version; \
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force $$version

migration-status:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version

# -------------------
# DATABASE
# -------------------
db-clean:
	@echo "🧨 Dropping database $(DB_NAME)..."
	psql -U $(DB_USER) -h $(DB_HOST) -c "DROP DATABASE IF EXISTS $(DB_NAME);"

db:
	@echo "⚠️ Dropping and recreating database $(DB_NAME)..."
	psql -U $(DB_USER) -h $(DB_HOST) -c "DROP DATABASE IF EXISTS $(DB_NAME);"
	psql -U $(DB_USER) -h $(DB_HOST) -c "CREATE DATABASE $(DB_NAME);"
	make migration-run
	make seed
