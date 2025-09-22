package pkg

import (
	"database/sql"
	"errors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"os"
)

func RunDb() (*sql.DB, error) {

	url := os.Getenv("DATABASE_URL")

	if url == "" {
		return nil, errors.New("DATABASE_URL environment variable is required")
	}

	con, err := sql.Open("pgx", url)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	log.Println("Database connected")

	return con, nil
}
