package dto

type ErrorResponse struct {
	Error   string                 `json:"error" example:"validation error"`
	Message string                 `json:"message,omitempty" example:"Invalid input data"`
	Details map[string]interface{} `json:"details,omitempty"`
}

type SuccessResponse struct {
	Success bool        `json:"success" example:"true"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty" example:"Operation completed successfully"`
}

type TokenResponse struct {
	Success bool        `json:"success" example:"true"`
	Token   string      `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User    interface{} `json:"user"`
}

type PaginatedResponse struct {
	Success bool        `json:"success" example:"true"`
	Data    interface{} `json:"data"`
	Page    int         `json:"page" example:"1"`
	PerPage int         `json:"per_page" example:"10"`
	Total   int64       `json:"total" example:"100"`
}

type TodoResponse struct {
	ID     uint   `json:"id" example:"1"`
	Name   string `json:"name" example:"My Todo"`
	UserID uint   `json:"user_id" example:"1"`
}

type TodoListResponse struct {
	Todos []interface{} `json:"todos"`
	Count int           `json:"count" example:"5"`
}

type UserResponse struct {
	ID    uint   `json:"id" example:"1"`
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"john@example.com"`
}
