package service

import (
	"context"
	"errors"
	"testing"
	"todo-list/internal/dto"
	"todo-list/internal/model"
	"todo-list/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestTodoService_CreateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTodoRepositoryI(ctrl)
	service := NewTodoService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name         string
		request      dto.CreateTodoRequest
		mockBehavior func()
		wantErr      bool
	}{
		{
			name: "Success",
			request: dto.CreateTodoRequest{
				Name:   "Test Todo",
				UserID: 1,
			},
			mockBehavior: func() {
				mockRepo.EXPECT().
					CreateTodo(ctx, gomock.Any()).
					Return(&model.Todo{ID: 1, Name: "Test Todo", UserID: 1}, nil)
			},
			wantErr: false,
		},
		{
			name: "Repository Error",
			request: dto.CreateTodoRequest{
				Name:   "Test Todo",
				UserID: 1,
			},
			mockBehavior: func() {
				mockRepo.EXPECT().
					CreateTodo(ctx, gomock.Any()).
					Return(nil, errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			todo, err := service.CreateTodo(ctx, tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, todo)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, todo)
				assert.Equal(t, tt.request.Name, todo.Name)
			}
		})
	}
}

func TestTodoService_GetTodoById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTodoRepositoryI(ctrl)
	service := NewTodoService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name         string
		id           uint
		mockBehavior func()
		wantErr      bool
	}{
		{
			name: "Success",
			id:   1,
			mockBehavior: func() {
				mockRepo.EXPECT().
					GetTodoById(ctx, uint(1)).
					Return(&model.Todo{ID: 1, Name: "Test Todo"}, nil)
			},
			wantErr: false,
		},
		{
			name: "Not Found",
			id:   999,
			mockBehavior: func() {
				mockRepo.EXPECT().
					GetTodoById(ctx, uint(999)).
					Return(nil, errors.New("not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			todo, err := service.GetTodoById(ctx, tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, todo)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, todo)
				assert.Equal(t, uint64(tt.id), todo.ID)
			}
		})
	}
}

func TestTodoService_UpdateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTodoRepositoryI(ctrl)
	service := NewTodoService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name         string
		id           uint
		payload      dto.UpdateTodoRequest
		mockBehavior func()
		wantErr      bool
	}{
		{
			name: "Success",
			id:   1,
			payload: dto.UpdateTodoRequest{
				Name:   "Updated Name",
				UserID: 1,
			},
			mockBehavior: func() {
				mockRepo.EXPECT().
					UpdateTodo(ctx, gomock.Any()).
					Return(&model.Todo{ID: 1, Name: "Updated Name", UserID: 1}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			todo, err := service.UpdateTodo(ctx, tt.payload, tt.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.payload.Name, todo.Name)
			}
		})
	}
}

func TestTodoService_DeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTodoRepositoryI(ctrl)
	service := NewTodoService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name         string
		id           uint
		userID       uint
		mockBehavior func()
		wantErr      bool
		expected     bool
	}{
		{
			name:   "Success",
			id:     1,
			userID: 1,
			mockBehavior: func() {
				mockRepo.EXPECT().
					DeleteTodo(ctx, uint(1), uint(1)).
					Return(true, nil)
			},
			wantErr:  false,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			deleted, err := service.DeleteTodo(ctx, tt.id, tt.userID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, deleted)
			}
		})
	}
}

func TestTodoService_ListTodos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTodoRepositoryI(ctrl)
	service := NewTodoService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name          string
		userID        uint
		page          int
		limit         int
		mockBehavior  func()
		wantErr       bool
		expectedLen   int
		expectedTotal int64
	}{
		{
			name:   "Success - Page 1",
			userID: 1,
			page:   1,
			limit:  10,
			mockBehavior: func() {
				mockRepo.EXPECT().
					ListTodos(ctx, uint(1), 10, 0).
					Return([]*model.Todo{{ID: 1}, {ID: 2}}, int64(2), nil)
			},
			wantErr:       false,
			expectedLen:   2,
			expectedTotal: 2,
		},
		{
			name:   "Success - Page 2",
			userID: 1,
			page:   2,
			limit:  5,
			mockBehavior: func() {
				mockRepo.EXPECT().
					ListTodos(ctx, uint(1), 5, 5).
					Return([]*model.Todo{{ID: 6}}, int64(10), nil)
			},
			wantErr:       false,
			expectedLen:   1,
			expectedTotal: 10,
		},
		{
			name:   "Repository Error",
			userID: 1,
			page:   1,
			limit:  10,
			mockBehavior: func() {
				mockRepo.EXPECT().
					ListTodos(ctx, uint(1), 10, 0).
					Return(nil, int64(0), errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			todos, total, err := service.ListTodos(ctx, tt.userID, tt.page, tt.limit)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedTotal, total)
				assert.Len(t, todos, tt.expectedLen)
			}
		})
	}
}
