package service

import (
	"context"
	"errors"
	"testing"
	"time"
	"todo-list/internal/model"
	"todo-list/internal/repository/mocks"
	"todo-list/pkg"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryI(ctrl)
	jwtManager, _ := pkg.NewJWTManager("test_secret", time.Hour)
	service := NewUserService(mockRepo, jwtManager)
	ctx := context.Background()

	tests := []struct {
		name         string
		user         *model.User
		mockBehavior func()
		wantErr      bool
		expected     bool
	}{
		{
			name: "Success",
			user: &model.User{
				Name:     "John Doe",
				Email:    "john@example.com",
				Password: "password123",
			},
			mockBehavior: func() {
				mockRepo.EXPECT().
					CreateUser(ctx, gomock.Any()).
					Return(true, nil)
			},
			wantErr:  false,
			expected: true,
		},
		{
			name: "Repository Error",
			user: &model.User{
				Name:     "John Doe",
				Email:    "john@example.com",
				Password: "password123",
			},
			mockBehavior: func() {
				mockRepo.EXPECT().
					CreateUser(ctx, gomock.Any()).
					Return(false, errors.New("db error"))
			},
			wantErr:  true,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			created, err := service.CreateUser(ctx, tt.user)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, created)
			}
		})
	}
}

func TestUserService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryI(ctrl)
	jwtManager, _ := pkg.NewJWTManager("test_secret", time.Hour)
	service := NewUserService(mockRepo, jwtManager)
	ctx := context.Background()

	hashedPassword, _ := pkg.HashPassword("password123")
	dbUser := &model.User{
		ID:       1,
		Email:    "john@example.com",
		Password: hashedPassword,
	}

	tests := []struct {
		name         string
		payload      *model.User
		mockBehavior func()
		wantErr      bool
		errMsg       string
	}{
		{
			name: "Success",
			payload: &model.User{
				Email:    "john@example.com",
				Password: "password123",
			},
			mockBehavior: func() {
				mockRepo.EXPECT().
					GetUserByEmail(ctx, "john@example.com").
					Return(dbUser, nil)
			},
			wantErr: false,
		},
		{
			name: "User Not Found",
			payload: &model.User{
				Email:    "notfound@example.com",
				Password: "password123",
			},
			mockBehavior: func() {
				mockRepo.EXPECT().
					GetUserByEmail(ctx, "notfound@example.com").
					Return(nil, errors.New("not found"))
			},
			wantErr: true,
			errMsg:  "invalid credentials",
		},
		{
			name: "Wrong Password",
			payload: &model.User{
				Email:    "john@example.com",
				Password: "wrongpassword",
			},
			mockBehavior: func() {
				mockRepo.EXPECT().
					GetUserByEmail(ctx, "john@example.com").
					Return(dbUser, nil)
			},
			wantErr: true,
			errMsg:  "invalid credentials",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			token, user, err := service.Login(ctx, tt.payload)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
				assert.Empty(t, token)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
				assert.NotNil(t, user)
				assert.Equal(t, dbUser.Email, user.Email)
			}
		})
	}
}

func TestUserService_GetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryI(ctrl)
	service := NewUserService(mockRepo, nil)
	ctx := context.Background()

	tests := []struct {
		name         string
		email        string
		mockBehavior func()
		wantErr      bool
	}{
		{
			name:  "Success",
			email: "john@example.com",
			mockBehavior: func() {
				mockRepo.EXPECT().
					GetUserByEmail(ctx, "john@example.com").
					Return(&model.User{Email: "john@example.com"}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			user, err := service.GetUserByEmail(ctx, tt.email)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.email, user.Email)
			}
		})
	}
}
