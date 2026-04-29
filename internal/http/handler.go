package http

import (
	"errors"
	"net/http"
	"todo-list/internal/config"
	"todo-list/internal/dto"
	"todo-list/internal/model"
	"todo-list/internal/service"
	"todo-list/pkg"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *service.Services
}

func newHandler(cfg *config.Config) *Handler {
	jwtManager, err := pkg.NewJWTManager(cfg.JWT.Secret, cfg.JWT.TokenDuration)
	if err != nil {
		pkg.Logger.Fatal().Err(err).Msg("Failed to initialize JWT Manager")
	}

	return &Handler{
		Services: service.NewServices(jwtManager),
	}
}

// CreateTodo godoc
// @Summary Create a new todo
// @Description Create a new todo item for the authenticated user
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body dto.CreateTodoRequest true "Todo data"
// @Success 201 {object} model.Todo
// @Failure 400 {object} map[string]interface{}
// @Security BearerAuth
// @Router /todos/create [post]
func (h *Handler) CreateTodo(c *gin.Context) {
	var createRequest dto.CreateTodoRequest

	if err := c.Bind(&createRequest); err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()
	createRequest.UserID = c.GetUint("uid")

	todo, err := h.Services.TodoService.CreateTodo(ctx, createRequest)
	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// UpdateTodo godoc
// @Summary Update a todo
// @Description Update an existing todo by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body dto.UpdateTodoRequest true "Updated todo data"
// @Success 200 {object} model.Todo
// @Failure 400 {object} map[string]interface{}
// @Security BearerAuth
// @Router /todos/{id} [put]
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

	ctx := c.Request.Context()
	data.UserID = c.GetUint("uid")

	todo, err := h.Services.TodoService.UpdateTodo(ctx, data, params.ID)
	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	pkg.SuccessResponse(c, todo)
}

// DeleteTodo godoc
// @Summary Delete a todo
// @Description Delete a todo by ID
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Security BearerAuth
// @Router /todos/{id} [delete]
func (h *Handler) DeleteTodo(c *gin.Context) {
	var params DeleteTodoParams

	if err := c.ShouldBindUri(&params); err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()
	deleted, err := h.Services.TodoService.DeleteTodo(ctx, params.ID, c.GetUint("uid"))
	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": deleted,
	})
}

// ListTodos godoc
// @Summary List all todos
// @Description Get all todos for the authenticated user
// @Tags todos
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /todos/ [get]
func (h *Handler) ListTodos(c *gin.Context) {
	var pg dto.PaginationRequest
	if err := c.ShouldBindQuery(&pg); err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()
	todos, total, err := h.Services.TodoService.ListTodos(ctx, c.GetUint("uid"), pg.Page, pg.Limit)
	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dto.PaginatedResponse{
		Success: true,
		Data:    todos,
		Page:    pg.Page,
		PerPage: pg.Limit,
		Total:   total,
	})
}

// CreateTask godoc
// @Summary Create a task for a todo
// @Description Create a new task for a specific todo
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param task body dto.CreateTaskTodoRequest true "Task data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Security BearerAuth
// @Router /task/{id} [post]
func (h *Handler) CreateTask(c *gin.Context) {
	var params CreateTaskParams

	if err := c.ShouldBindUri(&params); err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()
	todo, err := h.Services.TodoService.GetTodoById(ctx, params.ID)
	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusInternalServerError)
		return
	}

	if todo == nil {
		pkg.ErrorResponse(c, errors.New("todo not found"), http.StatusNotFound)
		return
	}

	var data dto.CreateTaskTodoRequest

	if err = c.Bind(&data); err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest) // исправлено: 400, не 500
		return
	}

	_, err = h.Services.TaskService.CreateTask(ctx, data, params.ID)
	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created",
	})
}

// UpdateStatusTask godoc
// @Summary Update task status
// @Description Update the status of a specific task
// @Tags tasks
// @Accept json
// @Produce json
// @Param todoId path int true "Todo ID"
// @Param taskId path int true "Task ID"
// @Param status body dto.UpdateStatusTaskTodoRequest true "Task status"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Security BearerAuth
// @Router /task/{todoId}/{taskId} [put]
func (h *Handler) UpdateStatusTask(c *gin.Context) {
	var params UpdateStatusTaskParams

	if err := c.ShouldBindUri(&params); err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest) // исправлено: BindUri до Bind тела
		return
	}

	var data dto.UpdateStatusTaskTodoRequest

	if err := c.Bind(&data); err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest) // исправлено: 400, не 500
		return
	}

	ctx := c.Request.Context()
	_, err := h.Services.TodoService.GetTodoById(ctx, params.TodoId)
	if err != nil {
		pkg.ErrorResponse(c, errors.New("todo not found"), http.StatusNotFound)
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

// RegisterUser godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.RegisterUser true "User registration data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /register [post]
func (h *Handler) RegisterUser(c *gin.Context) {
	var request dto.RegisterUser

	if err := c.Bind(&request); err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	user := &model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	ctx := c.Request.Context()
	created, err := h.Services.UserService.CreateUser(ctx, user)
	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	if !created {
		pkg.ErrorResponse(c, errors.New("user already exists"), http.StatusConflict) // 409, не 400
		return
	}

	pkg.CreatedResponse(c, "User created")
}

// LoginUser godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body dto.LoginUser true "Login credentials"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /login [post]
func (h *Handler) LoginUser(c *gin.Context) {
	var request dto.LoginUser

	if err := c.Bind(&request); err != nil {
		pkg.ErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()
	payload := &model.User{
		Email:    request.Email,
		Password: request.Password,
	}

	token, user, err := h.Services.UserService.Login(ctx, payload)
	if err != nil {
		pkg.ErrorResponse(c, err, http.StatusUnauthorized) // исправлено: 401, не 404
		return
	}

	pkg.TokenResponse(c, user, token)
}
