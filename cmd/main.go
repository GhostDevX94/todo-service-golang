package main

import (
	"log"
	"todo-list/internal/config"
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

	if err := godotenv.Load(); err != nil {
		log.Println("[WARN] .env file not found, using system environment variables")
	}

	cfg, err := config.Load()
	if err != nil {
		panic("config error: " + err.Error())
	}

	pkg.InitLogger(cfg.Log.Level, cfg.App.GinMode)

	pkg.Logger.Info().Msg("Starting Todo List API")

	pkg.Logger.Info().
		Str("dir", cfg.DB.MigrationsDir).
		Msg("Running database migrations")

	if err := pkg.RunMigrations(cfg.DB.URL, cfg.DB.MigrationsDir); err != nil {
		pkg.Logger.Fatal().Err(err).Msg("Failed to run migrations")
	}

	route := http.NewRoute(cfg)
	route.RouteRun()

}
