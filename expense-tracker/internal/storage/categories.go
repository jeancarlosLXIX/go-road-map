package storage

import (
	"encoding/json"
	"expense-tracker/internal/models"

	// "expense-tracker/internal/utils"
	// "expense-tracker/internal/utils"
	"os"
)

type CategoryStore struct {
	FilePath string
}

func (this *CategoryStore) List() ([]models.Category, error) {

	var categories []models.Category
	file, err := os.ReadFile(this.FilePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &categories)

	return categories, nil
}
