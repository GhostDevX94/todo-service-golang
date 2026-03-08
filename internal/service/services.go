package service

import (
	"time"
	"todo-list/internal/repository"
	"todo-list/pkg"
)

var ctxTime = 10 * time.Second

type Services struct {
	UserService *UserService
	TodoService *TodoService
	TaskService *TaskService
	JWTManager  *pkg.JWTManager
}

func NewServices(jwtManager *pkg.JWTManager) *Services {
	newRepository := repository.NewRepository()
	return &Services{
		UserService: NewUserService(newRepository.UserRepository, jwtManager),
		TodoService: NewTodoService(newRepository.TodoRepository),
		TaskService: NewTaskService(newRepository.TaskRepository),
		JWTManager:  jwtManager,
	}
}
