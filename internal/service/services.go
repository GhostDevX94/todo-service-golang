package service

import "todo-list/internal/repository"

type Services struct {
	UserService *UserService
	TodoService *TodoService
	TaskService *TaskService
}

func NewServices() *Services {
	newRepository := repository.NewRepository()
	return &Services{
		UserService: NewUserService(newRepository),
		TodoService: NewTodoService(newRepository),
		TaskService: NewTaskService(newRepository),
	}
}
