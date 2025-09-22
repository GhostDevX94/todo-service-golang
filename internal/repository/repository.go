package repository

import (
	"log"
	"todo-list/pkg"
)

type Repository struct {
	UserRepository *UserRepository
	TodoRepository *TodoRepository
	TaskRepository *TaskRepository
}

func NewRepository() *Repository {

	db, err := pkg.RunDb()

	if err != nil {
		log.Fatalf("Error database connect: %v", err)
	}

	return &Repository{
		UserRepository: newUserRepository(db),
		TodoRepository: newTodoRepository(db),
		TaskRepository: newTaskRepository(db),
	}
}
