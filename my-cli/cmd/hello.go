/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/user"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "This command will greed you.",
	Run: func(cmd *cobra.Command, args []string) {
		getName, _ := cmd.Flags().GetString("name")

		fmt.Println("Hello,", getName)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
	currentUser, err := user.Current()

	if err != nil {
		helloCmd.PersistentFlags().StringP("name", "n", "nameless", "A help for foo")
	} else {
		helloCmd.PersistentFlags().StringP("name", "n", currentUser.Username, "A help for foo")

	}
	// helloCmd.PersistentFlags().StringVarP(&variable,...)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

}
