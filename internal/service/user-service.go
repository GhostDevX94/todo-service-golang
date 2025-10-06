package service

import (
	"context"
	"errors"
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

func (u *UserService) CreateUser(ctx context.Context, user *model.User) (bool, error) {
	
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

func (u *UserService) Login(ctx context.Context, payload *model.User) (string, *model.User, error) {
	user, err := u.repo.UserRepository.GetUserByEmail(ctx, payload.Email)

	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	checkPassword := pkg.CheckPasswordHash(payload.Password, user.Password)

	if !checkPassword {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := pkg.CreateJWTToken(user)

	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (u *UserService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	User, err := u.repo.UserRepository.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return User, nil
}
