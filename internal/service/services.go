package service

import (
	"todo-list/internal/repository"
	"todo-list/pkg"
)

type Services struct {
	UserService UserServiceI
	TodoService TodoServiceI
	TaskService TaskServiceI
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
