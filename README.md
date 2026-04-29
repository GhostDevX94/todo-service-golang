# 📝 Todo List API

[![Go Version](https://img.shields.io/badge/Go-1.24-00ADD8?style=flat&logo=go)](https://golang.org)
[![API Docs](https://img.shields.io/badge/API-Swagger-85EA2D?style=flat&logo=swagger)](http://localhost:8080/swagger/index.html)

A clean and efficient RESTful API for managing todos and tasks, built with Go (Gin), PostgreSQL, and JWT authentication.

## ✨ Features
- 🔐 **Auth:** Secure JWT authentication (Login/Register).
- 📑 **Docs:** Interactive Swagger UI documentation.
- 🐳 **Docker:** Full Docker & Docker Compose support.
- 📈 **Pagination:** Built-in pagination for todo lists.
- ✅ **Tests:** High coverage using Table-Driven tests (Service Layer 86%+).

## 🚀 Quick Start

### 1. Using Docker (Recommended)
```bash
make docker-up
```
*API: http://localhost:8080 | Swagger: /swagger/index.html*

### 2. Local Development
```bash
cp env.example .env     # Configure your database
make migrate-up         # Run migrations
make run                # Start the server
```

## 🧪 Testing
The project follows Go best practices with Table-Driven tests and Mocking.
```bash
make test               # Run all tests
go test ./... -cover    # Check coverage
```

## 🏗️ Project Structure
- `cmd/main.go` - Application entry point.
- `internal/http/` - Handlers, routes, and middlewares.
- `internal/service/` - Business logic layer (fully tested).
- `internal/repository/` - Data access layer (PostgreSQL).
- `pkg/` - Shared packages (JWT, Logger, Response helpers).
- `migrations/` - SQL migration files.

## ⚙️ Configuration
Key environment variables in `.env`:
- `DATABASE_URL`: PostgreSQL connection string.
- `JWT_SECRET`: Secret key for signing tokens.
- `APP_PORT`: HTTP server port (default 8080).

---
**Built with ❤️ using Go and Gin Framework**
