package repository

import "database/sql"

type UserRepository struct {
	db *sql.DB
}

func newUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
