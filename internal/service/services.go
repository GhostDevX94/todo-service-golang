package service

import (
	"time"
	"todo-list/internal/repository"
)

var ctxTime = 10 * time.Second

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
