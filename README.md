# 📝 Todo List API

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![API Docs](https://img.shields.io/badge/API-Swagger-85EA2D?style=flat&logo=swagger)](http://localhost:8080/swagger/index.html)

A production-ready RESTful API for managing todos and tasks, built with Go, Gin framework, and PostgreSQL. Features JWT authentication, structured logging, rate limiting, and comprehensive API documentation.

## ✨ Features

- 🔐 **JWT Authentication** - Secure user authentication and authorization
- 📊 **RESTful API** - Clean and intuitive API design
- 🗄️ **PostgreSQL Database** - Reliable data persistence with migrations
- 📝 **Swagger Documentation** - Interactive API documentation
- 🔍 **Structured Logging** - Production-ready logging with zerolog
- ⚡ **Rate Limiting** - Protection against abuse
- 🐳 **Docker Support** - Easy deployment with Docker and Docker Compose
- ✅ **Input Validation** - Comprehensive request validation
- 🛡️ **Error Handling** - Structured error responses

## 🚀 Quick Start

### Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd todo-list
   ```

2. **Start services with Docker Compose**
   ```bash
   make docker-up
   ```

3. **Access the API**
   - API: http://localhost:8080
   - Swagger UI: http://localhost:8080/swagger/index.html
   - Health Check: http://localhost:8080/health

4. **Stop services**
   ```bash
   make docker-down
   ```

### Local Development

1. **Prerequisites**
   - Go 1.21 or higher
   - PostgreSQL 15+
   - Make

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp env.example .env
   # Edit .env with your configuration
   ```

4. **Run database migrations**
   ```bash
   make migrate-up
   ```

5. **Start the server**
   ```bash
   make run
   ```

## 📚 API Documentation

### Authentication

#### Register a new user
```bash
POST /register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securePassword123"
}
```

#### Login
```bash
POST /login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "securePassword123"
}
```

### Todos (Requires Authentication)

All todo endpoints require the `Authorization: Bearer <token>` header.

#### List all todos
```bash
GET /todos/
Authorization: Bearer <your-jwt-token>
```

#### Create a todo
```bash
POST /todos/create
Authorization: Bearer <your-jwt-token>
Content-Type: application/json

{
  "name": "Buy groceries"
}
```

#### Update a todo
```bash
PUT /todos/{id}
Authorization: Bearer <your-jwt-token>
Content-Type: application/json

{
  "name": "Buy groceries and cook dinner"
}
```

#### Delete a todo
```bash
DELETE /todos/{id}
Authorization: Bearer <your-jwt-token>
```

### Tasks (Requires Authentication)

#### Create a task for a todo
```bash
POST /task/{todoId}
Authorization: Bearer <your-jwt-token>
Content-Type: application/json

{
  "title": "Buy milk"
}
```

#### Update task status
```bash
PUT /task/{todoId}/{taskId}
Authorization: Bearer <your-jwt-token>
Content-Type: application/json

{
  "status": true
}
```

For complete API documentation with examples, visit the [Swagger UI](http://localhost:8080/swagger/index.html) when the server is running.

## 🏗️ Project Structure

```
.
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── dto/                 # Data Transfer Objects
│   │   ├── request.go       # Request DTOs with validation
│   │   └── response.go      # Response DTOs
│   ├── errors/              # Custom error types
│   │   └── errors.go
│   ├── http/                # HTTP layer
│   │   ├── handler.go       # Request handlers
│   │   ├── middleware.go    # Middleware (auth, logging, rate limiting)
│   │   └── router.go        # Route definitions
│   ├── model/               # Database models
│   │   ├── User.go
│   │   ├── Todo.go
│   │   └── Task.go
│   ├── repository/          # Data access layer
│   │   ├── repository.go
│   │   ├── user.go
│   │   ├── todo.go
│   │   └── task.go
│   └── service/             # Business logic layer
│       ├── services.go
│       ├── user-service.go
│       ├── todo-service.go
│       └── task-service.go
├── pkg/                     # Shared packages
│   ├── db.go               # Database connection
│   ├── hash.go             # Password hashing
│   ├── jwt.go              # JWT utilities
│   ├── logger.go           # Structured logging
│   └── response.go         # Response helpers
├── migrations/              # Database migrations
├── docs/                    # Swagger documentation (auto-generated)
├── Dockerfile              # Docker image definition
├── docker-compose.yml      # Docker Compose configuration
├── Makefile                # Development commands
└── README.md               # This file
```

## ⚙️ Configuration

Configuration is done via environment variables. See `env.example` for all available options.

### Key Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | PostgreSQL connection string | Required |
| `APP_PORT` | HTTP server port | `8080` |
| `GIN_MODE` | Gin mode (debug/release) | `release` |
| `LOG_LEVEL` | Logging level (debug/info/warn/error) | `info` |
| `JWT_SECRET` | Secret key for JWT tokens | Required |
| `CORS_ALLOWED_ORIGINS` | Allowed CORS origins | `*` |

## 🛠️ Development

### Available Make Commands

```bash
# Build and run
make build          # Build the application
make run            # Build and run the application
make stop           # Stop the running application
make restart        # Restart the application

# Database migrations
make create-migration    # Create a new migration
make migrate-up         # Apply all migrations
make migrate-down       # Rollback last migration

# Docker
make docker-build       # Build Docker image
make docker-up          # Start with Docker Compose
make docker-down        # Stop Docker Compose
make docker-logs        # View application logs
make docker-rebuild     # Rebuild and restart

# Code quality
make lint              # Run linter
make fmt               # Format code
make test              # Run tests

# Documentation
make swagger           # Regenerate Swagger docs
```

### Regenerating Swagger Documentation

After modifying API endpoints or documentation comments:

```bash
make swagger
```

## 🧪 Testing

Run tests with coverage:

```bash
make test
```

## 📦 Deployment

### Docker Deployment

1. Build the image:
   ```bash
   docker build -t todo-list-api:latest .
   ```

2. Run with Docker Compose:
   ```bash
   docker-compose up -d
   ```

### Production Considerations

- Set `GIN_MODE=release`
- Use strong `JWT_SECRET`
- Configure `CORS_ALLOWED_ORIGINS` appropriately
- Set `LOG_LEVEL=info` or `warn`
- Use environment-specific database credentials
- Enable HTTPS/TLS in production
- Consider using a reverse proxy (nginx, Caddy)

## 🏥 Health Checks

The API provides a health check endpoint:

```bash
GET /health
```

Response:
```json
{
  "status": "healthy",
  "service": "todo-list-api"
}
```

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📧 Contact

For questions or support, please contact: support@todolist.com

---

**Built with ❤️ using Go and Gin Framework**
