/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"expense-tracker/internal/storage"

	"github.com/spf13/cobra"
)

// sumaryCmd represents the sumary command
var sumaryCmd = &cobra.Command{
	Use:   "sumary",
	Short: "Print the total amount spent this month.",
	Run: func(cmd *cobra.Command, args []string) {
		ExpenseObj := storage.ExpenseStore{FilePath: "expenses.json"}

		ExpenseObj.Sumary()
	},
}

func init() {
	rootCmd.AddCommand(sumaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sumaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sumaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
