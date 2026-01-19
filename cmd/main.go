package main

import (
	"todo-list/internal/http"
	"todo-list/pkg"

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

	pkg.InitLogger()

	pkg.Logger.Info().Msg("Starting Todo List API")

	if err := godotenv.Load(); err != nil {
		pkg.Logger.Warn().Err(err).Msg(".env file not found, using environment variables")
	}

	route := http.NewRoute()
	route.RouteRun()

}
