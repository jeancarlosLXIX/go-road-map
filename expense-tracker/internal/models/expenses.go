package models

import "time"

type Expense struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	CategoryId int       `json:"category_id"`
	Total      float64   `json:"total"`
	CreatedAt  time.Time `json:"created_at"`
}
