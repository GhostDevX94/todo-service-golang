# Todo List API

REST API для управления списками задач с пользователями.

## Структура проекта

```
todo-list/
├── cmd/
│   ├── main.go          # Основное приложение
│   └── migrate/         # Утилита для миграций
│       └── main.go
├── internal/
│   ├── http/            # HTTP handlers и роутинг
│   ├── model/           # Модели данных
│   ├── repository/      # Слой доступа к данным
│   └── service/         # Бизнес-логика
├── migrations/          # SQL миграции
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Установка и запуск

### Предварительные требования

- Go 1.24+
- PostgreSQL 12+

### Установка зависимостей

```bash
go mod download
go mod tidy
```

### Настройка базы данных

1. Создайте базу данных PostgreSQL:
```sql
CREATE DATABASE todo_list;
```

2. Настройте переменную окружения с URL базы данных:
```bash
export DB_URL="postgres://username:password@localhost:5432/todo_list?sslmode=disable"
```

## Миграции

Проект использует `golang-migrate/migrate` для управления схемой базы данных.

### Доступные миграции

1. **000001_create_users_table** - Создание таблицы пользователей
2. **000002_create_todos_table** - Создание таблицы списков задач
3. **000003_create_tak_todos_table** - Создание таблицы задач

### Управление миграциями

#### Через Makefile (рекомендуется)

```bash
# Применить все миграции
make migrate-up

# Применить определенное количество миграций
make migrate-up MIGRATE_STEPS=2

# Откатить все миграции
make migrate-down

# Откатить определенное количество миграций
make migrate-down MIGRATE_STEPS=1

# Показать текущую версию
make migrate-version

# Принудительно установить версию
make migrate-force MIGRATE_VERSION=3

# Показать справку по командам
make help
```

#### Напрямую через Go

```bash
# Применить все миграции
go run cmd/migrate/main.go -action=up -db="postgres://user:pass@localhost:5432/dbname?sslmode=disable"

# Применить 2 миграции
go run cmd/migrate/main.go -action=up -steps=2 -db="postgres://user:pass@localhost:5432/dbname?sslmode=disable"

# Откатить 1 миграцию
go run cmd/migrate/main.go -action=down -steps=1 -db="postgres://user:pass@localhost:5432/dbname?sslmode=disable"

# Показать версию
go run cmd/migrate/main.go -action=version -db="postgres://user:pass@localhost:5432/dbname?sslmode=disable"
```

### Схема базы данных

#### Таблица users
- `id` - Уникальный идентификатор (SERIAL PRIMARY KEY)
- `name` - Имя пользователя (VARCHAR(255))
- `email` - Email пользователя (VARCHAR(255), UNIQUE)
- `password` - Хеш пароля (VARCHAR(255))
- `created_at` - Время создания (TIMESTAMP WITH TIME ZONE)
- `updated_at` - Время обновления (TIMESTAMP WITH TIME ZONE)

#### Таблица todos
- `id` - Уникальный идентификатор (SERIAL PRIMARY KEY)
- `name` - Название списка задач (VARCHAR(255))
- `user_id` - ID пользователя (INTEGER, FOREIGN KEY)
- `date` - Дата (TIMESTAMP WITH TIME ZONE)
- `created_at` - Время создания (TIMESTAMP WITH TIME ZONE)
- `updated_at` - Время обновления (TIMESTAMP WITH TIME ZONE)

#### Таблица tak_todos
- `id` - Уникальный идентификатор (SERIAL PRIMARY KEY)
- `title` - Название задачи (VARCHAR(255))
- `todo_id` - ID списка задач (INTEGER, FOREIGN KEY)
- `status` - Статус выполнения (INTEGER, 0=не выполнено, 1=выполнено)
- `date` - Дата (TIMESTAMP WITH TIME ZONE)
- `created_at` - Время создания (TIMESTAMP WITH TIME ZONE)
- `updated_at` - Время обновления (TIMESTAMP WITH TIME ZONE)

### Индексы

Для оптимизации производительности созданы следующие индексы:

- `idx_users_email` - Поиск пользователей по email
- `idx_users_name` - Поиск пользователей по имени
- `idx_todos_user_id` - Поиск списков задач по пользователю
- `idx_todos_date` - Поиск списков задач по дате
- `idx_todos_user_date` - Составной индекс по пользователю и дате
- `idx_tak_todos_todo_id` - Поиск задач по списку
- `idx_tak_todos_status` - Поиск задач по статусу
- `idx_tak_todos_date` - Поиск задач по дате
- `idx_tak_todos_todo_status` - Составной индекс по списку и статусу

## Сборка и запуск

### Сборка

```bash
make build
```

### Запуск

```bash
make run
```

### Очистка

```bash
make clean
```

## Тестирование

```bash
make test
```

## Разработка

### Добавление новой миграции

1. Создайте файлы `migrations/XXXXXX_description.up.sql` и `migrations/XXXXXX_description.down.sql`
2. Номер миграции должен быть больше текущего максимального
3. В `up.sql` опишите изменения для применения
4. В `down.sql` опишите изменения для отката

### Пример новой миграции

```sql
-- migrations/000004_add_priority_to_tasks.up.sql
ALTER TABLE tak_todos ADD COLUMN priority INTEGER DEFAULT 1;

-- migrations/000004_add_priority_to_tasks.down.sql
ALTER TABLE tak_todos DROP COLUMN priority;
```

## Лицензия

MIT
