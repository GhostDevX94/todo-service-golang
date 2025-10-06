package repository

import (
	"context"
	"database/sql"
	"todo-list/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func newUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserById(ctx context.Context, id uint) (*model.User, error) {
	return nil, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, data *model.User) (bool, error) {

	row := r.db.QueryRowContext(ctx, "INSERT INTO users (name,email,password) VALUES ($1,$2,$3) RETURNING id,name,email,password", data.Name, data.Email, data.Password)

	err := row.Scan(
		&data.ID,
		&data.Name,
		&data.Email,
		&data.Password,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {

	var user model.User

	row := r.db.QueryRowContext(ctx, "SELECT id,name,email,password FROM users WHERE email = $1", email)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
