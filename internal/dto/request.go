package dto

type CreateTodoRequest struct {
	Name   string `json:"name" binding:"required,min=1,max=255"`
	UserID uint   `json:"-"`
}

type UpdateTodoRequest struct {
	Name   string `json:"name" binding:"required,min=1,max=255"`
	UserID uint   `json:"-"`
}

type CreateTaskTodoRequest struct {
	Title string `json:"title"`
}

type UpdateStatusTaskTodoRequest struct {
	Status bool `json:"status"`
}

type RegisterUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
