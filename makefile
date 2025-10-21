.DEFAULT_GOAL := help

# ==========================
# ⚙️ Docker commands
# ==========================

build: ## Build docker containers
	@echo "🔨 Building..."
	docker compose build --no-cache

run: ## Start containers
	@echo "🚀 Running app..."
	docker compose up --build

down: ## Stop and delete containers
	@echo "🧹 Stopping and removing containers..."
	docker compose down

logs: ## Docker containers logs
	@echo "📜 Showing logs..."
	docker compose logs -f app

# ==========================
# 🗃️ Database migrations
# ==========================

DB_URL = postgres://postgres:322434@db:5432/go?sslmode=disable
MIGRATIONS_PATH = internal/db/migrations

migrate-up: ## Apply all migrations
	@echo "📈 Applying migrations..."
	docker compose exec app migrate -path /app/migrations -database "$(DB_URL)" up

migrate-down: ## Roll back last migration
	@echo "📉 Rolling back last migration..."
	docker compose exec app migrate -path /app/migrations -database "$(DB_URL)" down 1

migrate-create: ## Create new migration
ifeq ($(strip $(name)),)
	@echo "❌ Enter migration name: make migrate-create name=user"
else
	@echo "🆕 Creating migration '$(name)'..."
	migrate create -ext sql -dir $(MIGRATIONS_PATH) $(name)
endif

# ==========================
# 📘 Help
# ==========================
help:
	@echo "Available commands:"
	@echo "  make build - build project"
	@echo "  make run - start project"
	@echo "  make down - delete project"
	@echo "  make logs - project logs"
	@echo "  make migrate-up - apply migrations"
	@echo "  make migrate-down - roll back last migrations"
	@echo "  make migrate-create name=NAME - create new migrations"
