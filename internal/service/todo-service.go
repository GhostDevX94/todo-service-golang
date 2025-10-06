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
	ListTodos(context.Context, uint) ([]*model.Todo, error)
}

type TodoService struct {
	repo *repository.Repository
}

func NewTodoService(repo *repository.Repository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (s *TodoService) GetTodoById(ctx context.Context, id uint) (*model.Todo, error) {
	todo, err := s.repo.TodoRepository.GetTodoById(ctx, id)
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

	todo, err := s.repo.TodoRepository.CreateTodo(ctx, todo)

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

	todo, err := s.repo.TodoRepository.UpdateTodo(ctx, payloadTodo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, id uint, UserId uint) (bool, error) {
	deleted, err := s.repo.TodoRepository.DeleteTodo(ctx, id, UserId)

	if err != nil {
		return false, err
	}

	return deleted, nil
}

func (s *TodoService) ListTodos(ctx context.Context, UserId uint) ([]*model.Todo, error) {
	todos, err := s.repo.TodoRepository.ListTodos(ctx, UserId)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
