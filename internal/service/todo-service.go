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
	DeleteTodo(context.Context, uint) (bool, error)
	ListTodos(context.Context) ([]*model.Todo, error)
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
	todo, err := s.repo.TodoRepository.CreateTodo(ctx, request)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, request dto.UpdateTodoRequest, id uint) (*model.Todo, error) {
	todo, err := s.repo.TodoRepository.UpdateTodo(ctx, request, id)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, id uint) (bool, error) {
	deleted, err := s.repo.TodoRepository.DeleteTodo(ctx, id)

	if err != nil {
		return false, err
	}

	return deleted, nil
}

func (s *TodoService) ListTodos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := s.repo.TodoRepository.ListTodos(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
