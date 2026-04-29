package service

import (
	"context"
	"todo-list/internal/dto"
	"todo-list/internal/model"
	"todo-list/internal/repository"
)

type TodoServiceI interface {
	CreateTodo(context.Context, dto.CreateTodoRequest) (*model.Todo, error)
	GetTodoById(context.Context, uint) (*model.Todo, error)
	UpdateTodo(context.Context, dto.UpdateTodoRequest, uint) (*model.Todo, error)
	DeleteTodo(context.Context, uint, uint) (bool, error)
	ListTodos(context.Context, uint, int, int) ([]*model.Todo, int64, error)
}

type TodoService struct {
	repo repository.TodoRepositoryI
}

func NewTodoService(repo repository.TodoRepositoryI) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (s *TodoService) GetTodoById(ctx context.Context, id uint) (*model.Todo, error) {
	todo, err := s.repo.GetTodoById(ctx, id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) CreateTodo(ctx context.Context, request dto.CreateTodoRequest) (*model.Todo, error) {

	todo := &model.Todo{
		Name:   request.Name,
		UserID: request.UserID,
	}

	todo, err := s.repo.CreateTodo(ctx, todo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, payload dto.UpdateTodoRequest, id uint) (*model.Todo, error) {

	payloadTodo := &model.Todo{
		ID:     uint64(id),
		Name:   payload.Name,
		UserID: payload.UserID,
	}

	todo, err := s.repo.UpdateTodo(ctx, payloadTodo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, id uint, UserId uint) (bool, error) {
	deleted, err := s.repo.DeleteTodo(ctx, id, UserId)

	if err != nil {
		return false, err
	}

	return deleted, nil
}

func (s *TodoService) ListTodos(ctx context.Context, UserId uint, page, limit int) ([]*model.Todo, int64, error) {
	offset := (page - 1) * limit
	todos, total, err := s.repo.ListTodos(ctx, UserId, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return todos, total, nil
}
