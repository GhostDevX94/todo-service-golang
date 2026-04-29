package dto

type CreateTodoRequest struct {
	Name   string `json:"name" binding:"required,min=1,max=255" example:"Buy groceries"`
	UserID uint   `json:"-"`
}

type UpdateTodoRequest struct {
	Name   string `json:"name" binding:"required,min=1,max=255" example:"Updated todo name"`
	UserID uint   `json:"-"`
}

type CreateTaskTodoRequest struct {
	Title string `json:"title" binding:"required,min=1,max=255" example:"Complete task 1"`
}

type UpdateStatusTaskTodoRequest struct {
	Status bool `json:"status" binding:"required" example:"true"`
}

type RegisterUser struct {
	Name     string `json:"name" binding:"required,min=2,max=100" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required,min=6,max=100" example:"securePassword123"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"securePassword123"`
}

type PaginationRequest struct {
	Page  int `form:"page,default=1" binding:"omitempty,min=1"`
	Limit int `form:"limit,default=10" binding:"omitempty,min=1,max=100"`
}
