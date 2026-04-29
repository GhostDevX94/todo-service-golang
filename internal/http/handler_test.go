package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-list/internal/dto"
	"todo-list/internal/model"
	"todo-list/internal/service"
	"todo-list/internal/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func setupTestRouter(h *Handler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	return r
}

func TestHandler_RegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserSvc := mocks.NewMockUserServiceI(ctrl)
	h := &Handler{
		Services: &service.Services{
			UserService: mockUserSvc,
		},
	}

	router := setupTestRouter(h)
	router.POST("/register", h.RegisterUser)

	tests := []struct {
		name         string
		payload      interface{}
		mockBehavior func()
		expectedCode int
	}{
		{
			name: "Success",
			payload: dto.RegisterUser{
				Name:     "John",
				Email:    "john@example.com",
				Password: "password123",
			},
			mockBehavior: func() {
				mockUserSvc.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Return(true, nil)
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "Invalid Input",
			payload: map[string]string{
				"email": "invalid-email",
			},
			mockBehavior: func() {},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "User Already Exists",
			payload: dto.RegisterUser{
				Name:     "John",
				Email:    "john@example.com",
				Password: "password123",
			},
			mockBehavior: func() {
				mockUserSvc.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Return(false, nil)
			},
			expectedCode: http.StatusConflict,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			
			body, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}

func TestHandler_LoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserSvc := mocks.NewMockUserServiceI(ctrl)
	h := &Handler{
		Services: &service.Services{
			UserService: mockUserSvc,
		},
	}

	router := setupTestRouter(h)
	router.POST("/login", h.LoginUser)

	tests := []struct {
		name         string
		payload      interface{}
		mockBehavior func()
		expectedCode int
	}{
		{
			name: "Success",
			payload: dto.LoginUser{
				Email:    "john@example.com",
				Password: "password123",
			},
			mockBehavior: func() {
				mockUserSvc.EXPECT().
					Login(gomock.Any(), gomock.Any()).
					Return("test_token", &model.User{ID: 1, Email: "john@example.com"}, nil)
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "Unauthorized",
			payload: dto.LoginUser{
				Email:    "john@example.com",
				Password: "wrong_password",
			},
			mockBehavior: func() {
				mockUserSvc.EXPECT().
					Login(gomock.Any(), gomock.Any()).
					Return("", nil, assert.AnError)
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			body, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}

func TestHandler_CreateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoSvc := mocks.NewMockTodoServiceI(ctrl)
	h := &Handler{
		Services: &service.Services{
			TodoService: mockTodoSvc,
		},
	}

	router := setupTestRouter(h)
	// სიმულაცია AuthMiddleware-ის, რომელიც ადებს uid-ს კონტექსტში
	router.POST("/todos", func(c *gin.Context) {
		c.Set("uid", uint(1))
		h.CreateTodo(c)
	})

	tests := []struct {
		name         string
		payload      interface{}
		mockBehavior func()
		expectedCode int
	}{
		{
			name: "Success",
			payload: dto.CreateTodoRequest{
				Name: "New Todo",
			},
			mockBehavior: func() {
				mockTodoSvc.EXPECT().
					CreateTodo(gomock.Any(), gomock.Any()).
					Return(&model.Todo{ID: 1, Name: "New Todo"}, nil)
			},
			expectedCode: http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			body, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}

func TestHandler_ListTodos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoSvc := mocks.NewMockTodoServiceI(ctrl)
	h := &Handler{
		Services: &service.Services{
			TodoService: mockTodoSvc,
		},
	}

	router := setupTestRouter(h)
	router.GET("/todos", func(c *gin.Context) {
		c.Set("uid", uint(1))
		h.ListTodos(c)
	})

	tests := []struct {
		name         string
		query        string
		mockBehavior func()
		expectedCode int
	}{
		{
			name:  "Success",
			query: "?page=1&limit=10",
			mockBehavior: func() {
				mockTodoSvc.EXPECT().
					ListTodos(gomock.Any(), uint(1), 1, 10).
					Return([]*model.Todo{{ID: 1}}, int64(1), nil)
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodGet, "/todos"+tt.query, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}

func TestHandler_UpdateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoSvc := mocks.NewMockTodoServiceI(ctrl)
	h := &Handler{
		Services: &service.Services{
			TodoService: mockTodoSvc,
		},
	}

	router := setupTestRouter(h)
	router.PUT("/todos/:id", func(c *gin.Context) {
		c.Set("uid", uint(1))
		h.UpdateTodo(c)
	})

	tests := []struct {
		name         string
		id           string
		payload      interface{}
		mockBehavior func()
		expectedCode int
	}{
		{
			name: "Success",
			id:   "1",
			payload: dto.UpdateTodoRequest{
				Name: "Updated Todo",
			},
			mockBehavior: func() {
				mockTodoSvc.EXPECT().
					UpdateTodo(gomock.Any(), gomock.Any(), uint(1)).
					Return(&model.Todo{ID: 1, Name: "Updated Todo"}, nil)
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			body, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest(http.MethodPut, "/todos/"+tt.id, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}

func TestHandler_DeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoSvc := mocks.NewMockTodoServiceI(ctrl)
	h := &Handler{
		Services: &service.Services{
			TodoService: mockTodoSvc,
		},
	}

	router := setupTestRouter(h)
	router.DELETE("/todos/:id", func(c *gin.Context) {
		c.Set("uid", uint(1))
		h.DeleteTodo(c)
	})

	tests := []struct {
		name         string
		id           string
		mockBehavior func()
		expectedCode int
	}{
		{
			name: "Success",
			id:   "1",
			mockBehavior: func() {
				mockTodoSvc.EXPECT().
					DeleteTodo(gomock.Any(), uint(1), uint(1)).
					Return(true, nil)
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			req, _ := http.NewRequest(http.MethodDelete, "/todos/"+tt.id, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}

func TestHandler_CreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoSvc := mocks.NewMockTodoServiceI(ctrl)
	mockTaskSvc := mocks.NewMockTaskServiceI(ctrl)
	h := &Handler{
		Services: &service.Services{
			TodoService: mockTodoSvc,
			TaskService: mockTaskSvc,
		},
	}

	router := setupTestRouter(h)
	router.POST("/task/:id", h.CreateTask)

	tests := []struct {
		name         string
		id           string
		payload      interface{}
		mockBehavior func()
		expectedCode int
	}{
		{
			name: "Success",
			id:   "1",
			payload: dto.CreateTaskTodoRequest{
				Title: "New Task",
			},
			mockBehavior: func() {
				mockTodoSvc.EXPECT().
					GetTodoById(gomock.Any(), uint(1)).
					Return(&model.Todo{ID: 1}, nil)
				mockTaskSvc.EXPECT().
					CreateTask(gomock.Any(), gomock.Any(), uint(1)).
					Return(&model.TaskTodos{ID: 1}, nil)
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "Todo Not Found",
			id:   "999",
			payload: dto.CreateTaskTodoRequest{
				Title: "New Task",
			},
			mockBehavior: func() {
				mockTodoSvc.EXPECT().
					GetTodoById(gomock.Any(), uint(999)).
					Return(nil, nil)
			},
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			body, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest(http.MethodPost, "/task/"+tt.id, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}

func TestHandler_UpdateStatusTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoSvc := mocks.NewMockTodoServiceI(ctrl)
	mockTaskSvc := mocks.NewMockTaskServiceI(ctrl)
	h := &Handler{
		Services: &service.Services{
			TodoService: mockTodoSvc,
			TaskService: mockTaskSvc,
		},
	}

	router := setupTestRouter(h)
	router.PUT("/task/:todoId/:taskId", h.UpdateStatusTask)

	tests := []struct {
		name         string
		todoID       string
		taskID       string
		payload      interface{}
		mockBehavior func()
		expectedCode int
	}{
		{
			name:   "Success",
			todoID: "1",
			taskID: "1",
			payload: dto.UpdateStatusTaskTodoRequest{
				Status: true,
			},
			mockBehavior: func() {
				mockTodoSvc.EXPECT().
					GetTodoById(gomock.Any(), uint(1)).
					Return(&model.Todo{ID: 1}, nil)
				mockTaskSvc.EXPECT().
					UpdateStatusTask(gomock.Any(), gomock.Any(), uint(1), uint(1)).
					Return(true, nil)
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			body, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest(http.MethodPut, "/task/"+tt.todoID+"/"+tt.taskID, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}
