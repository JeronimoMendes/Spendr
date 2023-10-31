/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/JeronimoMendes/spendr/cmd"
	"github.com/spf13/cobra"
)

// expenseCmd represents the expense command
var expenseCmd = &cobra.Command{
	Use:   "expense",
	Run: func(cmd *cobra.Command, args []string) {
		// print the help message for this command
		fmt.Println(cmd.Help())
	},
}

func init() {
	cmd.RootCmd.AddCommand(expenseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// expenseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// expenseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
