/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"math"

	"github.com/JeronimoMendes/spendr/cmd"
	"github.com/JeronimoMendes/spendr/pkg/gc_client"
	"github.com/JeronimoMendes/spendr/pkg/tracker"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Show information about category",
	Run: func(cmd *cobra.Command, args []string) {
		secretID := viper.GetString("gc_secret_id")
		secretKey := viper.GetString("gc_secret_key")

		gc := gc_client.NewClient(secretKey, secretID)
		tracker := tracker.NewExpenseTracker(gc)

		var categories []string
		if len(args) > 0 {
			categories = args
		} else {
			for _, category := range tracker.Categories {
				categories = append(categories, category.Name)
			}
		}
		for _, categoryName := range categories {
			category, err := tracker.GetCategory(categoryName)
			if err != nil {
				fmt.Println(err)
				continue
			}

			totalAmount := math.Abs(tracker.GetTotalExpensesByCategory(categoryName))

			red := color.New(color.FgRed).SprintFunc()
			green := color.New(color.FgGreen).SprintFunc()
	
			fmt.Printf("[%s]\n", categoryName)
			if category.Limit > 0 {
				remainingAmount := category.Limit - totalAmount
				var remainingAmountColored string
				if remainingAmount < 0 {
					remainingAmountColored = red(fmt.Sprintf("%.2f", remainingAmount))
				} else {
					remainingAmountColored = green(fmt.Sprintf("%.2f", remainingAmount))
				}

				fmt.Printf("Total amount: €%.2f (%.2f%%)\n", totalAmount, totalAmount/category.Limit*100)
				fmt.Printf("Limit: €%.2f\n", category.Limit)
				fmt.Printf("Remaining amount: €%s\n\n", remainingAmountColored)
			} else {
				fmt.Printf("Total amount: €%.2f\n\n", totalAmount)
			}
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(categoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
