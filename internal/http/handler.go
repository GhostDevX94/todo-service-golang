package http

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-list/internal/dto"
	"todo-list/internal/model"
	"todo-list/internal/service"
	"todo-list/pkg"
)

var ctx = context.Background()

type Handler struct {
	Services *service.Services
}

func newHandler() *Handler {
	context.WithTimeout(ctx, 1000000)
	return &Handler{
		Services: service.NewServices(),
	}
}

func (h *Handler) CreateTodo(c *gin.Context) {

	var createRequest dto.CreateTodoRequest

	err := c.Bind(&createRequest)

	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}

	todo, err := h.Services.TodoService.CreateTodo(ctx, createRequest)

	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}

	c.JSON(http.StatusCreated, todo)
}

type UpdateTodoParams struct {
	ID uint `uri:"id" binding:"required"`
}

func (h *Handler) UpdateTodo(c *gin.Context) {
	var params UpdateTodoParams

	if err := c.ShouldBindUri(&params); err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	var data dto.UpdateTodoRequest

	if err := c.Bind(&data); err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	todo, err := h.Services.TodoService.UpdateTodo(ctx, data, params.ID)
	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	pkg.SuccessResponse(c, todo)
}

type DeleteTodoParams struct {
	ID uint `uri:"id" binding:"required"`
}

func (h *Handler) DeleteTodo(c *gin.Context) {
	var params DeleteTodoParams

	if err := c.ShouldBindUri(&params); err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	deleted, err := h.Services.TodoService.DeleteTodo(ctx, params.ID)
	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": deleted,
	})
}

func (h *Handler) ListTodos(c *gin.Context) {
	todos, err := h.Services.TodoService.ListTodos(ctx)
	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
		"count": len(todos),
	})
}

type CreateTaskParams struct {
	ID uint `uri:"id" binding:"required"`
}

func (h *Handler) CreateTask(c *gin.Context) {
	var params CreateTaskParams

	if err := c.ShouldBindUri(&params); err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	todo, err := h.Services.TodoService.GetTodoById(ctx, params.ID)
	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusInternalServerError)
		return
	}

	if todo == nil {
		pkg.ErrorResponse(c, errors.New("not found"), http.StatusNotFound)
		return
	}

	var data dto.CreateTaskTodoRequest

	_, err = h.Services.TaskService.CreateTask(ctx, data, params.ID)
	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task created",
	})

}

type UpdateStatusTaskParams struct {
	TodoId uint `uri:"todoId" binding:"required"`
	TaskID uint `uri:"taskId" binding:"required"`
}

func (h *Handler) UpdateStatusTask(c *gin.Context) {

	var data dto.UpdateStatusTaskTodoRequest

	err := c.Bind(&data)
	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusInternalServerError)
		return
	}
	var params UpdateStatusTaskParams

	err = c.ShouldBindUri(&params)
	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusInternalServerError)
		return
	}

	_, err = h.Services.TodoService.GetTodoById(ctx, params.TodoId)
	if err != nil {
		pkg.ErrorResponse(c, errors.New("not found"), http.StatusNotFound)
		return
	}

	updated, err := h.Services.TaskService.UpdateStatusTask(ctx, data, params.TodoId, params.TaskID)
	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": updated,
	})

}

func (h *Handler) RegisterUser(c *gin.Context) {

	var request dto.RegisterUser

	err := c.Bind(&request)

	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	user := &model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	created, err := h.Services.UserService.CreateUser(user)

	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	if created == false {
		pkg.ErrorResponse(c, errors.New("user already exists"), http.StatusBadRequest)
		return
	}

	pkg.CreatedResponse(c, "User created")

}

func (h *Handler) LoginUser(c *gin.Context) {
	var request dto.LoginUser

	err := c.Bind(&request)

	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	payload := &model.User{
		Email:    request.Email,
		Password: request.Password,
	}

	token, user, err := h.Services.UserService.Login(payload)

	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusNotFound)
		return
	}

	pkg.TokenResponse(c, user, token)

}
