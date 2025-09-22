package dto

type CreateTodoRequest struct {
	Name string `json:"name" binding:"required,min=1,max=255"`
}

type UpdateTodoRequest struct {
	Name string `json:"name" binding:"required,min=1,max=255"`
}

type CreateTaskTodoRequest struct {
	Title string `json:"title"`
}

type UpdateStatusTaskTodoRequest struct {
	Status bool `json:"status"`
}
