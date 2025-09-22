package model

import "time"

type Todo struct {
	ID     uint64    `json:"id"`
	Name   string    `json:"name"`
	UserID uint      `json:"user_id"`
	Date   time.Time `json:"date"`
}
