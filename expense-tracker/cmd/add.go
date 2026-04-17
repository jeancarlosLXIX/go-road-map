/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"expense-tracker/internal/models"
	"expense-tracker/internal/storage"
	"expense-tracker/internal/utils"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var total float64
var catID int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add expenses in your history",
	Example: `
To just add the expense:
add food

the default amount is 0

to add an amount
add food --total 100
add food -t 100

`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		CategoryObj := storage.CategoryStore{FilePath: "categories.json"}
		utils.DefaultCategory(CategoryObj.FilePath)

		expName := args[0]
		ExpenseObj := storage.ExpenseStore{FilePath: "expenses.json"}

		if ExpenseObj.ExpenseExists(expName) {
			fmt.Printf("%s already exists.\n", expName)
			return
		}

		if !utils.GetCategoryId(catID, CategoryObj.FilePath) {
			fmt.Println("Category with that id does not exist, adding default ID 0")
			catID = 0
		}

		exp := models.Expense{
			ID:         ExpenseObj.GetAnId(),
			Name:       expName,
			Total:      total,
			CategoryId: catID,
			CreatedAt:  time.Now(),
		}

		errorE := ExpenseObj.Add(exp)

		if errorE != nil {
			fmt.Println("Expense added correctly.")
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().Float64VarP(&total, "total", "t", 0, "Expense amount")
	addCmd.Flags().IntVarP(&catID, "category", "c", 0, "Category ID")
}
