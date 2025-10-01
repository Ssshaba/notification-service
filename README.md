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

- GET `http://localhost:8876/api/documentation` - получить уведомления
- POST `http://localhost:8876/api/documentation` - Создать сообщение синхронно
- POST `http://localhost:8876/api/documentation` - Создать сообщение через очередь

Пример JSON для создания уведомления
```json
{
  "sender": "sender@example.com",
  "recipient": "recipient@example.com",
  "message": "Текст уведомления"
}
```

Также можно создать уведомление через интерфейс RabbitMQ Management:

 1. Откройте http://localhost:15672, введите логин и пароль(если оставить дефолтные значения из примера в `.env` то логин и пароль = `guest`)
 2. Откройте: Queues and Streams->выбрать очередь notifications->Publish message
 3. Вставьте и отправьте пример json сообщения 