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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a category",
	Run: func(cmd *cobra.Command, args []string) {
		secretID := viper.GetString("gc_secret_id")
		secretKey := viper.GetString("gc_secret_key")

		gc := gc_client.NewClient(secretKey, secretID)
		tracker := tracker.NewExpenseTracker(gc)

		category := args[0]
		tracker.DeleteCategory(category)
		fmt.Printf("Category %s deleted.\n", category)
	},
}

func init() {
	categoryCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
