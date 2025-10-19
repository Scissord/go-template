# ==========================
# 1️⃣ Этап сборки
# ==========================
FROM golang:1.25.1 AS builder

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники
COPY . .

# Устанавливаем утилиту migrate (для работы с миграциями)
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Собираем бинарник
RUN go build -o main ./cmd/api


# ==========================
# 2️⃣ Этап запуска
# ==========================
FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Копируем бинарники
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/main .
COPY .env .env

# ⚙️ добавляем миграции в контейнер
COPY internal/db/migrations ./migrations

EXPOSE 8080

CMD ["./main"]
