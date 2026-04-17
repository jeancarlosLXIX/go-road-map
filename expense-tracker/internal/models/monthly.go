package models

type Monthly struct {
	Id              int       `json:"id"`
	Year            int       `json:"year"`
	Month           string    `json:"month"`
	MonthlyExpenses []Expense `json:"monthly_expenses"`
}
