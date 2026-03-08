package repository

import (
	"log"
	"todo-list/pkg"
)

type Repository struct {
	UserRepository UserRepositoryI
	TodoRepository TodoRepositoryI
	TaskRepository TaskRepositoryI
}

func NewRepository() *Repository {

	db, err := pkg.ConnectDB()

	if err != nil {
		log.Fatalf("Error database connect: %v", err)
	}

	return &Repository{
		UserRepository: newUserRepository(db),
		TodoRepository: newTodoRepository(db),
		TaskRepository: newTaskRepository(db),
	}
}
