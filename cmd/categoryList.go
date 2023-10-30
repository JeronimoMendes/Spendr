/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/JeronimoMendes/expense-track/pkg/gc_client"
	"github.com/JeronimoMendes/expense-track/pkg/tracker"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var categoryListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all categories",
	Run: func(cmd *cobra.Command, args []string) {
		secretID := viper.GetString("gc_secret_id")
		secretKey := viper.GetString("gc_secret_key")

		gc := gc_client.NewClient(secretKey, secretID)
		tracker := tracker.NewExpenseTracker(gc)

		categories := tracker.GetCategories()
		for _, category := range categories {
			println(category)
		}
	},
}

func init() {
	categoryCmd.AddCommand(categoryListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
