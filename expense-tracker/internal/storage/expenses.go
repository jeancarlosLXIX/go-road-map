package storage

import (
	"encoding/json"
	"expense-tracker/internal/models"
	"expense-tracker/internal/utils"
	"slices"
	"strings"

	// "expense-tracker/internal/utils"
	"os"
)

type ExpenseStore struct {
	FilePath string
}

// Add a expense to the file in FilePath
func (this *ExpenseStore) Add(exp models.Expense) error {

	var expenses []models.Expense

	// Read file (ignore error if file doesn't exist yet)
	file, _ := os.ReadFile(this.FilePath)

	// Load existing data
	if len(file) > 0 {
		json.Unmarshal(file, &expenses)
	}

	expenses = append(expenses, exp)
	// Save back to file
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(this.FilePath, data, 0644)
}

// Get all the values in our JSON file
func (this *ExpenseStore) GetAll() ([]models.Expense, error) {

	err := utils.FileExist(this.FilePath) // Check if the file exists

	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(this.FilePath)
	var expenses []models.Expense
	err = json.Unmarshal(file, &expenses)

	return expenses, err
}

// To check if a expense already exists in our file
func (this *ExpenseStore) ExpenseExists(name string) bool {

	expenses, _ := this.GetAll()

	for _, exp := range expenses {
		if strings.ToLower(name) == exp.Name {
			return true
		}
	}

	return false
}

// Return an id
func (this *ExpenseStore) GetAnId() int {
	expenses, _ := this.GetAll()

	// idSlice := []int{0}
	var idSlice []int

	// Filling the slices ID
	for _, v := range expenses {

		// if slices.Contains(idSlice, v.ID) {
		// 	continue
		// }
		idSlice = append(idSlice, v.ID)
	}

	// Getting an ID
	for i := 1; i <= 1254; i++ {
		if !slices.Contains(idSlice, i) {
			return i
		}
	}
	return 0
}

// listCmd.Flags().BoolVarP(&showCategories, "categories", "c", false, "Show categories instead of expenses")
