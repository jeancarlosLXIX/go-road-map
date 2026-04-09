package storage

import (
	"encoding/json"
	"os"

	"expense-tracker/internal/models"
)

type JSONStore struct {
	FilePath string
}

func (s *JSONStore) Load() ([]models.Expense, error) {
	file, err := os.ReadFile(s.FilePath)
	if err != nil {
		return nil, err
	}

	var expenses []models.Expense
	err = json.Unmarshal(file, &expenses)
	return expenses, err
}

func (s *JSONStore) Save(expenses []models.Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.FilePath, data, 0644)
}
