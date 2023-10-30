/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/JeronimoMendes/spendr/pkg/gc_client"
	"github.com/JeronimoMendes/spendr/pkg/tracker"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new category",
	Run: func(cmd *cobra.Command, args []string) {
		secretID := viper.GetString("gc_secret_id")
		secretKey := viper.GetString("gc_secret_key")

		gc := gc_client.NewClient(secretKey, secretID)
		tracker_client := tracker.NewExpenseTracker(gc)

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			panic(err)
		}
		limit, err := cmd.Flags().GetFloat64("limit")
		if err != nil {
			limit = -1
		}

		tracker_client.CreateCategory(tracker.Category{Name: name, Limit: limit})
		fmt.Printf("Category %s created.\n", name)
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

	createCmd.Flags().StringP("name", "n", "", "Name of the category")
	createCmd.Flags().Float64P("limit", "l", 0, "Limit of the category")
}
