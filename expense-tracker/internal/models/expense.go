package models

import "time"

type Expense struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Total      float64   `json:"total"`
	CategoryId int       `json:"category"`
	CreatedAt  time.Time `json:"created_at"`
}
