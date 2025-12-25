package main

import (
	"log"
	"todo-list/internal/http"

	"github.com/joho/godotenv"
)

// @title Todo List API
// @version 1.0
// @description RESTful API for managing todos and tasks
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@todolist.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {

	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	route := http.NewRoute()
	route.RouteRun()

}
