package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"todo-list/internal/dto"
	"todo-list/internal/model"
)

type TodoRepositoryI interface {
	CreateTodo(context.Context, dto.CreateTodoRequest) (*model.Todo, error)
	GetTodoById(context.Context, uint) (*model.Todo, error)
	UpdateTodo(context.Context, dto.UpdateTodoRequest, uint) (*model.Todo, error)
	DeleteTodo(context.Context, uint) (bool, error)
	ListTodos(context.Context) ([]*model.Todo, error)
}

type TodoRepository struct {
	db *sql.DB
}

func newTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) GetTodoById(ctx context.Context, id uint) (*model.Todo, error) {
	var todo model.Todo

	query := r.db.QueryRowContext(ctx, "SELECT id,name,user_id,date FROM todos WHERE id = $1", id)

	err := query.Scan(&todo.ID, &todo.Name, &todo.UserID, &todo.Date)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepository) CreateTodo(ctx context.Context, request dto.CreateTodoRequest) (*model.Todo, error) {
	var todo model.Todo
	query := r.db.QueryRowContext(ctx, "INSERT INTO todos (name,user_id,date) VALUES ($1,$2,$3) RETURNING id,name,user_id,date", request.Name, 1, time.Now())

	err := query.Scan(&todo.ID, &todo.Name, &todo.UserID, &todo.Date)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepository) UpdateTodo(ctx context.Context, request dto.UpdateTodoRequest, id uint) (*model.Todo, error) {
	var todo model.Todo
	query := r.db.QueryRowContext(ctx, "UPDATE todos SET name = $1  WHERE id = $2 RETURNING id,name,user_id,date", request.Name, id)

	err := query.Scan(&todo.ID, &todo.Name, &todo.UserID, &todo.Date)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepository) DeleteTodo(ctx context.Context, id uint) (bool, error) {
	result, err := r.db.ExecContext(ctx, "DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return false, fmt.Errorf("failed to delete todo: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("failed to get rows affected: %w", err)
	}

	return rowsAffected > 0, nil
}

func (r *TodoRepository) ListTodos(ctx context.Context) ([]*model.Todo, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, user_id, date FROM todos ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*model.Todo
	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.ID, &todo.Name, &todo.UserID, &todo.Date)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
