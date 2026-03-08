package pkg

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB() (*sql.DB, error) {

	url := os.Getenv("DATABASE_URL")

	if url == "" {
		return nil, errors.New("DATABASE_URL environment variable is required")
	}

	con, err := sql.Open("pgx", url)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := con.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	con.SetMaxOpenConns(25)
	con.SetMaxIdleConns(25)
	con.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Database connected successfully")

	return con, nil
}
