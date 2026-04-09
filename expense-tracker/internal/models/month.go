package models

type Month struct {
	Id       int       `json:"id"`
	Year     string    `json:"year"`
	Month    string    `json:"month"`
	Expenses []Expense `json:"expenses"`
}
