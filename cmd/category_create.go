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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new category",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		secretID := viper.GetString("gc_secret_id")
		secretKey := viper.GetString("gc_secret_key")

		gc := gc_client.NewClient(secretKey, secretID)
		tracker := tracker.NewExpenseTracker(gc)

		category := args[0]
		tracker.CreateCategory(category)
		fmt.Printf("Category %s created.\n", category)
	},
}

func init() {
	categoryCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
