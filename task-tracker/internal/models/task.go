package models

import "time"

type Task struct {
	ID          uint16    `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // TODO, IN-PROGRESS, DONE
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
