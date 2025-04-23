# auth-api

Простой сервис аутентификации на Go с использованием gRPC.

## Требования

- Go 1.23.4 или выше
- PostgreSQL 15 или выше
- Make (опционально)

## Установка

1. Клонируйте репозиторий:
```bash
git clone https://gitlab.crja72.ru/golang/2025/spring/course/projects/go21/auth-api.git
cd file-processing-api
```

2. Установите зависимости:
```bash
go mod download
```

3. Создайте файл конфигурации `.env` в корне проекта:
```bash
MIGRATIONS_PATH=/migrations

TOKEN_SECRET=SUPER_SECRET

# PostgreSQL конфигурация
PG_USER=postgres
PG_PASSWORD=postgres
PG_HOST=localhost
PG_PORT=5432
PG_DBNAME=postgres
DB_MAX_CONNS=2
DB_CONN_TIMEOUT=30s
```

## Запуск приложения

### Подготовка базы данных

1. Создайте базу данных PostgreSQL:
```bash
createdb postgres
```

2. Примените миграции:
```bash
make migrate-up
```

### Если вы используете Docker:
```bash
docker-compose up --build
```

### Запуск сервера

1. В режиме разработки:
```bash
make run-dev
```

2. В production режиме:
```bash
make run
```

## Доступные команды

- `make run` - запуск приложения
- `make run-dev` - запуск в режиме разработки
- `make migrate-up` - применение миграций
- `make migrate-down` - откат миграций
- `make migrations-create name=<name>` - создание новой миграции
- `make proto-generate` - генерация gRPC контрактов
- `make generate-key` - генерация секретного ключа

## Тестирование

Для запуска тестов:
```bash
go test ./...
```

Для проверки покрытия тестами:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

## CI/CD

Проект использует GitLab CI/CD для автоматизации сборки и тестирования. Пайплайн включает:
- Сборку приложения
- Запуск тестов
- Проверку покрытия кода тестами (минимум 30%)

## Структура проекта

```
├── cmd/                    # Точки входа приложения
│   ├── app/               # Основное приложение
│   ├── migrate/           # Миграции базы данных
│   └── generate-key/      # Генерация ключей
├── config/                # Конфигурация
├── gen/                   # Сгенерированные файлы
├── internal/              # Внутренние пакеты
│   ├── entity/            # Сущности
│   ├── app/               # Точка входа (откуда идёт запуск)
│   ├── controller/        # Контроллеры (transport)
│   ├── repository/        # Репозитории
│   └── usecase/           # Сценарии использования
├── migrations/            # SQL миграции
├── init-scripts/          # Скрипты для поднятия базы
├── pkg/                   # Публичные пакеты
│   ├── logger/            # Работа с логгером (zap)
│   └── validator/         # Валидация
└── proto/                 # gRPC протофайлы
```

## API

### gRPC методы

- `Register` - регистрация нового пользователя
- `Login` - авторизация пользователя
- `RefreshToken` - обновление токена
- `ValidateToken` - проверка токена
- `Logout` - выход из системы

