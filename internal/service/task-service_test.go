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

func TestTaskService_CreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTaskRepositoryI(ctrl)
	service := NewTaskService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name         string
		request      dto.CreateTaskTodoRequest
		todoID       uint
		mockBehavior func()
		wantErr      bool
	}{
		{
			name: "Success",
			request: dto.CreateTaskTodoRequest{
				Title: "New Task",
			},
			todoID: 1,
			mockBehavior: func() {
				mockRepo.EXPECT().
					CreateTask(ctx, gomock.Any(), uint(1)).
					Return(&model.TaskTodos{ID: 1, Title: "New Task"}, nil)
			},
			wantErr: false,
		},
		{
			name: "Repository Error",
			request: dto.CreateTaskTodoRequest{
				Title: "New Task",
			},
			todoID: 1,
			mockBehavior: func() {
				mockRepo.EXPECT().
					CreateTask(ctx, gomock.Any(), uint(1)).
					Return(nil, errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			task, err := service.CreateTask(ctx, tt.request, tt.todoID)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, task)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, task)
				assert.Equal(t, tt.request.Title, task.Title)
			}
		})
	}
}

func TestTaskService_UpdateStatusTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTaskRepositoryI(ctrl)
	service := NewTaskService(mockRepo)
	ctx := context.Background()

	tests := []struct {
		name         string
		data         dto.UpdateStatusTaskTodoRequest
		todoID       uint
		taskID       uint
		mockBehavior func()
		wantErr      bool
		expected     bool
	}{
		{
			name: "Success",
			data: dto.UpdateStatusTaskTodoRequest{
				Status: true,
			},
			todoID: 1,
			taskID: 1,
			mockBehavior: func() {
				mockRepo.EXPECT().
					UpdateStatusTask(ctx, gomock.Any(), uint(1), uint(1)).
					Return(true, nil)
			},
			wantErr:  false,
			expected: true,
		},
		{
			name: "Task Not Found",
			data: dto.UpdateStatusTaskTodoRequest{
				Status: true,
			},
			todoID: 1,
			taskID: 999,
			mockBehavior: func() {
				mockRepo.EXPECT().
					UpdateStatusTask(ctx, gomock.Any(), uint(1), uint(999)).
					Return(false, nil)
			},
			wantErr:  false,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			updated, err := service.UpdateStatusTask(ctx, tt.data, tt.todoID, tt.taskID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, updated)
			}
		})
	}
}
