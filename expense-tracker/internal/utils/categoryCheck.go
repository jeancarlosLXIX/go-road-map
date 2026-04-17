package utils

import (
	"encoding/json"
	"expense-tracker/internal/models"
	"os"
	"time"
)

// Create a fileName.json file for some defaults values in case
// the file doesn't exist
func DefaultCategory(fileName string) error {
	// Check if file exists
	_, err := os.Stat(fileName)

	if err == nil {
		return nil
	}

	// If error is NOT "file does not exist", return it
	if !os.IsNotExist(err) {
		return err
	}

	defaultCategory := []models.Category{
		{
			ID:        0,
			Name:      "others",
			CreatedAt: time.Now(),
		},
	}

	data, err := json.MarshalIndent(defaultCategory, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}

// Get the category ID ofr the value
func GetCategoryId(idToCheck int, categoryPath string) bool {

	// catName = strings.ToLower(catName)
	var categories []models.Category

	data, err := os.ReadFile(categoryPath)
	json.Unmarshal(data, &categories)

	for _, val := range categories {
		if val.ID == idToCheck {
			return true
		}
	}

	if err != nil {
		panic("Error to get ID")
	}

	return false

}
