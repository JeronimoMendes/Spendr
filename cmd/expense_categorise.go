/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/JeronimoMendes/expense-track/pkg/gc_client"
	"github.com/JeronimoMendes/expense-track/pkg/tracker"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// categoriseCmd represents the categorise command
var categoriseCmd = &cobra.Command{
	Use:   "categorise",
	Short: "Set the category for an expense",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		secretID := viper.GetString("gc_secret_id")
		secretKey := viper.GetString("gc_secret_key")

		gc := gc_client.NewClient(secretKey, secretID)
		tracker := tracker.NewExpenseTracker(gc)

		tracker.CategoriseExpense(args[0], args[1])
		fmt.Println("Expenses categorised.")
	},
}

func init() {
	expenseCmd.AddCommand(categoriseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoriseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoriseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
