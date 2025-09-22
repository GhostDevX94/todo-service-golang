package model

import "time"

type TaskTodos struct {
	ID     uint64    `json:"id"`
	Title  string    `json:"title"`
	TodoID uint      `json:"todo_id"`
	Status bool      `json:"status"`
	Date   time.Time `json:"date"`
}
