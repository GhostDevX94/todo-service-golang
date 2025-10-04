package service

import (
	"context"
	"todo-list/internal/model"
	"todo-list/internal/repository"
	"todo-list/pkg"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) CreateUser(user *model.User) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), ctxTime)
	defer cancel()

	password, err := pkg.HashPassword(user.Password)

	if err != nil {
		return false, err
	}

	user.Password = password

	created, err := u.repo.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return false, err
	}

	return created, nil
}

func (u *UserService) Login(payload *model.User) (string, *model.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), ctxTime)
	defer cancel()

	user, err := u.repo.UserRepository.GetUserByEmail(ctx, payload.Email)

	if err != nil {
		return "", nil, err
	}

	checkPassword := pkg.CheckPasswordHash(payload.Password, user.Password)

	if !checkPassword {
		return "", nil, err
	}

	token, err := pkg.CreateJWTToken(user)

	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (u *UserService) GetUserByEmail(email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTime)
	defer cancel()

	User, err := u.repo.UserRepository.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return User, nil
}
