package storage

import (
	"encoding/json"
	"os"

	"expense-tracker/internal/models"
)

type ExpenseStore struct {
	FilePath string
}

func (s *ExpenseStore) Load() ([]models.Expense, error) {
	if err := fileExist(s.FilePath); err != nil {
		return nil, err
	}

	data, err := os.ReadFile(s.FilePath)
	if err != nil {
		return nil, err
	}

	var expenses []models.Expense
	err = json.Unmarshal(data, &expenses)
	return expenses, err
}

func (s *ExpenseStore) Save(expenses []models.Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.FilePath, data, 0644)
}

func (s *ExpenseStore) PrintList(expenses []models.Expense) {
	// do something
}
