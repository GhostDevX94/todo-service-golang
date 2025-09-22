package repository

import (
	"context"
	"database/sql"
	"time"
	"todo-list/internal/dto"
	"todo-list/internal/model"
)

type TaskRepositoryI interface {
	CreateTask(context.Context, dto.CreateTaskTodoRequest, uint) (*model.TaskTodos, error)
	UpdateStatusTask(context.Context, dto.UpdateStatusTaskTodoRequest, uint, uint) (bool, error)
}

type TaskRepository struct {
	db *sql.DB
}

func newTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) CreateTask(ctx context.Context, request dto.CreateTaskTodoRequest, TodoId uint) (*model.TaskTodos, error) {
	var task model.TaskTodos

	query := r.db.QueryRowContext(ctx, "INSERT INTO task_todos (title,todo_id,date) VALUES ($1,$2,$3) RETURNING id,title,date", request.Title, TodoId, time.Now())

	err := query.Scan(&task.ID, &task.Title, &task.Date)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) UpdateStatusTask(ctx context.Context, data dto.UpdateStatusTaskTodoRequest, TodoId uint, TaskId uint) (bool, error) {

	query := r.db.QueryRowContext(ctx, "UPDATE task_todos SET status = $1 WHERE todo_id = $2 AND id = $3;", data.Status, TodoId, TaskId)

	if query.Err() != nil {
		return false, query.Err()
	}

	return true, nil
}
