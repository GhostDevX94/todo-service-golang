package config

import (
	"errors"
	"os"
	"time"
)

type Config struct {
	App AppConfig
	DB  DBConfig
	JWT JWTConfig
	Log LogConfig
}

type AppConfig struct {
	Port        string
	GinMode     string
	CORSOrigins string
}

type DBConfig struct {
	URL           string
	MigrationsDir string
}

type JWTConfig struct {
	Secret        string
	TokenDuration time.Duration
}

type LogConfig struct {
	Level string
}

func Load() (*Config, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET environment variable is required")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, errors.New("DATABASE_URL environment variable is required")
	}

	migrationsDir := os.Getenv("MIGRATIONS_DIR")
	if migrationsDir == "" {
		migrationsDir = "migrations"
	}

	tokenDuration := 24 * time.Hour
	if d := os.Getenv("JWT_DURATION"); d != "" {
		if parsed, err := time.ParseDuration(d); err == nil {
			tokenDuration = parsed
		}
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "release"
	}

	corsOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if corsOrigins == "" {
		corsOrigins = "*"
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	return &Config{
		App: AppConfig{
			Port:        port,
			GinMode:     ginMode,
			CORSOrigins: corsOrigins,
		},
		DB: DBConfig{
			URL:           dbURL,
			MigrationsDir: migrationsDir,
		},
		JWT: JWTConfig{
			Secret:        jwtSecret,
			TokenDuration: tokenDuration,
		},
		Log: LogConfig{
			Level: logLevel,
		},
	}, nil
}
