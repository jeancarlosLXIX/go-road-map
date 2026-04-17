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

var categoryID int
var name string

// var total float64

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test [name]",
	Short: "Add values to your expenses history",
	// Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		expenseObj := storage.ExpenseStore{FilePath: "expense.json"}
		CatetegoryObj := storage.CategoryStore{FilePath: "categories.json"}
		utils.DefaultCategory(CatetegoryObj.FilePath)
		// exp, err := expenseObj.GetAll()

		// if err != nil {
		// 	fmt.Println("Error getting the expenses.")
		// 	return
		// }

		// for _, e := range exp {
		// 	fmt.Printf("%s - %.2f\n", e.Name, e.Total)
		// }

		expense := models.Expense{
			ID:         1,
			Name:       "Default",
			Total:      0,
			CategoryId: 0,
			CreatedAt:  time.Now(),
		}
		// expenses, err := expenseObj.GetAll()
		err := expenseObj.Add(expense)

		if err != nil {
			fmt.Println("An error has occurred")
			return
		}
		// fmt.Println("ADDED.")
		// fmt.Println(expenses)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	testCmd.Flags().StringVarP(&name, "name", "n", "", "Expense name")
	testCmd.Flags().Float64VarP(&total, "total", "t", 0, "Expense total")
	testCmd.Flags().IntVarP(&categoryID, "category", "c", 0, "Category ID (default: 0)")

	testCmd.MarkFlagRequired("name")
}
