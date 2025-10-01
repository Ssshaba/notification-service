# Test Case Notification Service

Golang/Gin, PostgreSQL, RebbitMQ

## Требования

- `Docker`
- `Docker Compose`

## Подготовка

### 1. Создайте файл `.env`** в корне проекта.

### 2. Скопируйте в него содержимое файла `.env.example`
### 3. Придумайте свои значения для переменных или оставьте значения по умолчанию

## Установка и запуск

### 1. Собрать и поднять контейнеры
`docker compose -f docker-compose.yml up --build -d`
---
После этих шагов будут доступны следующие маршруты:

- GET `http://localhost:8876/api/documentation`
- POST `http://localhost:8876/api/documentation`
- POST `http://localhost:8876/api/documentation`

