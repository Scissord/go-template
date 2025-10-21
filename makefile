.DEFAULT_GOAL := help

# ==========================
# ⚙️ Docker commands
# ==========================

build: ## Собрать контейнер
	@echo "🔨 Building..."
	docker compose build --no-cache

run: ## Запустить приложение
	@echo "🚀 Running app..."
	docker compose up --build

down: ## Остановить и удалить контейнеры
	@echo "🧹 Stopping and removing containers..."
	docker compose down

logs: ## Посмотреть логи приложения
	@echo "📜 Showing logs..."
	docker compose logs -f app

# ==========================
# 🗃️ Database migrations
# ==========================
# Для работы нужно, чтобы в контейнере app был установлен migrate (или другая утилита)
# Пример: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

DB_URL = postgres://postgres:322434@db:5432/go?sslmode=disable
MIGRATIONS_PATH = internal/db/migrations

migrate-up: ## Применить все новые миграции
	@echo "📈 Applying migrations..."
	docker compose exec app migrate -path /app/migrations -database "$(DB_URL)" up

migrate-down: ## Откатить последнюю миграцию
	@echo "📉 Rolling back last migration..."
	docker compose exec app migrate -path /app/migrations -database "$(DB_URL)" down 1

migrate-create: ## Создать новую миграцию
ifeq ($(strip $(name)),)
	@echo "❌ Укажи имя миграции: make migrate-create name=users_table"
else
	@echo "🆕 Creating migration '$(name)'..."
	migrate create -ext sql -dir $(MIGRATIONS_PATH) $(name)
endif

# ==========================
# 📘 Help
# ==========================
help:
	@echo ""
	@echo "📘 Available commands:"
	@echo "  make build            — собрать проект"
	@echo "  make run              — запустить проект"
	@echo "  make down             — удалить проект"
	@echo "  make logs             — логи проекта"
	@echo "  make migrate-up       — применить все миграции"
	@echo "  make migrate-down     — откатить последнюю миграцию"
	@echo "  make migrate-create name=NAME — создать новую миграцию"
	@echo ""
