package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Making this to test",
	Long:  "Making this to test what can I do by myself",
	// Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"testo", "car"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("you must provide a name")
		}

		fmt.Println("Hello,", args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
