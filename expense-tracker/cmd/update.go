/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"strings"

	"github.com/spf13/cobra"

	// MINE
	"expense-tracker/internal/storage"
)

var amount float64

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [expense]",
	Short: "Update the expenses.",
	Example: `
Use:
update food --amount 100
update food -a 100
`,
	Run: func(cmd *cobra.Command, args []string) {
		if amount == 0 {
			fmt.Println("Please add an amount to update")
			return
		}
		ExpenseObj := storage.ExpenseStore{FilePath: "expenses.json"}
		expToUpdate := strings.ToLower(args[0])

		if ExpenseObj.ExpenseExists(expToUpdate) {
			ExpenseObj.Update(expToUpdate, amount)
		} else {
			fmt.Println("Expense " + expToUpdate + " Doesn't exist.")
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Amount to update")
}
