Todo List Service

This is a simple RESTful API for managing todos and their tasks. It is built with Go using the Gin framework and stores data in PostgreSQL (via `pgx`).

Available endpoints:
- GET `/todos/` — list all todos
- POST `/todos/create` — create a new todo
- PUT `/todos/:id` — update a todo
- DELETE `/todos/:id` — delete a todo
- POST `/task/:id` — create a task for the given todo
- PUT `/task/:todoId/:taskId` — update a task status for the given todo

Configuration via environment variables:
- `DATABASE_URL` (required): PostgreSQL connection string
- `APP_PORT` (optional, default `8080`): HTTP port
- `GIN_MODE` (optional, default `release`): Gin run mode

Database migrations are located in the `migrations` directory. Run the application from `cmd/main.go` or use the provided binary/script under `bin/`.

