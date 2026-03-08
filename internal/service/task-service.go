package service

import (
	"context"
	"todo-list/internal/dto"
	"todo-list/internal/model"
	"todo-list/internal/repository"
)

type TaskServiceI interface {
	CreateTask(context.Context, dto.CreateTaskTodoRequest, uint) (*model.TaskTodos, error)
	UpdateStatusTask(context.Context, dto.UpdateStatusTaskTodoRequest, uint, uint) (bool, error)
}

type TaskService struct {
	repo repository.TaskRepositoryI
}

func NewTaskService(repo repository.TaskRepositoryI) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (t *TaskService) CreateTask(ctx context.Context, request dto.CreateTaskTodoRequest, TodoId uint) (*model.TaskTodos, error) {

	task, err := t.repo.CreateTask(ctx, request, TodoId)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *TaskService) UpdateStatusTask(ctx context.Context, data dto.UpdateStatusTaskTodoRequest, TodoId uint, TaskId uint) (bool, error) {
	status, err := t.repo.UpdateStatusTask(ctx, data, TodoId, TaskId)
	if err != nil {
		return false, err
	}

	return status, nil
}
