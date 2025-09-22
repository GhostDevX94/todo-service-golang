package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-list/internal/dto"
	"todo-list/internal/service"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	todo, err := h.Services.TodoService.CreateTodo(ctx, createRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, todo)
}

type UpdateTodoParams struct {
	ID uint `uri:"id" binding:"required"`
}

func (h *Handler) UpdateTodo(c *gin.Context) {
	var params UpdateTodoParams

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var data dto.UpdateTodoRequest

	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.Services.TodoService.UpdateTodo(ctx, data, params.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

type DeleteTodoParams struct {
	ID uint `uri:"id" binding:"required"`
}

func (h *Handler) DeleteTodo(c *gin.Context) {
	var params DeleteTodoParams

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deleted, err := h.Services.TodoService.DeleteTodo(ctx, params.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": deleted,
	})
}

func (h *Handler) ListTodos(c *gin.Context) {
	todos, err := h.Services.TodoService.ListTodos(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.Services.TodoService.GetTodoById(ctx, params.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if todo == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not found",
		})
		return
	}

	var data dto.CreateTaskTodoRequest

	_, err = h.Services.TaskService.CreateTask(ctx, data, params.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var params UpdateStatusTaskParams

	err = c.ShouldBindUri(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = h.Services.TodoService.GetTodoById(ctx, params.TodoId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not found",
		})
		return
	}

	updated, err := h.Services.TaskService.UpdateStatusTask(ctx, data, params.TodoId, params.TaskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": updated,
	})

}
