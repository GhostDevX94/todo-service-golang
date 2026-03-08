package http

// URI-параметры для роутов (id из path)

type UpdateTodoParams struct {
	ID uint `uri:"id" binding:"required"`
}

type DeleteTodoParams struct {
	ID uint `uri:"id" binding:"required"`
}

type CreateTaskParams struct {
	ID uint `uri:"id" binding:"required"`
}

type UpdateStatusTaskParams struct {
	TodoId uint `uri:"todoId" binding:"required"`
	TaskID uint `uri:"taskId" binding:"required"`
}
