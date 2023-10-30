/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/JeronimoMendes/spendr/pkg/gc_client"
	"github.com/JeronimoMendes/spendr/pkg/tracker"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// categoriseCmd represents the categorise command
var categoriseCmd = &cobra.Command{
	Use:   "categorise",
	Short: "Set the category for an expense",
	Run: func(cmd *cobra.Command, args []string) {
		secretID := viper.GetString("gc_secret_id")
		secretKey := viper.GetString("gc_secret_key")

		gc := gc_client.NewClient(secretKey, secretID)
		tracker := tracker.NewExpenseTracker(gc)

		if len(args) == 0 {
			expenses := tracker.GetExpensesByCategory("")
			if len(expenses) == 0 {
				fmt.Println("No expenses to categorise.")
				return
			}
			for _, expense := range expenses {
				w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
				fmt.Fprintln(w, "ID\tDescription\tAmount\tDate")
				fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%.2f\t%s\t%s", expense.Id, expense.Description, expense.Amount, expense.Date, expense.Category))
				category := ""
				w.Flush()

				categoriesNames := []string{}
				for _, category := range tracker.Categories {
					categoriesNames = append(categoriesNames, category.Name)
				}
				prompt := promptui.Select{
					Label: "Select Category",
					Items: append([]string{"skip"}, categoriesNames...),
				}
				_, category, err := prompt.Run()
				if err != nil {
					fmt.Printf("Prompt failed %v\n", err)
					return
				}

				if category == "skip" {
					continue
				}
				tracker.CategoriseExpense(expense.Id, category)
			}
		} else {
			tracker.CategoriseExpense(args[0], args[1])
			fmt.Println("Expenses categorised.")
		}
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
