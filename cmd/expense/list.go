/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package expense

import (
	"fmt"

	"github.com/JeronimoMendes/spendr/pkg/gc_client"
	"github.com/JeronimoMendes/spendr/pkg/tracker"

	"github.com/spf13/cobra"

	"os"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var expenseListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all expenses",
	Run: func(cmd *cobra.Command, args []string) {
		secretID := viper.GetString("gc_secret_id")
		secretKey := viper.GetString("gc_secret_key")
		account_id := viper.GetString("gc_account_id")

		gc := gc_client.NewClient(secretKey, secretID)
		tracker := tracker.NewExpenseTracker(gc)

		full, _ := cmd.Flags().GetBool("full")
		update, _ := cmd.Flags().GetBool("update")
		transactions := tracker.GetExpenses(account_id, full, update)	

		red := color.New(color.FgRed).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
		defer w.Flush()
		fmt.Fprintln(w, "ID\tDescription\tAmount\tDate\tCategory")
		fmt.Fprintln(w, "--\t-----------\t------\t----\t--------")
		for _, transaction := range transactions {
			var amountColored string
			if transaction.Amount < 0 {
				amountColored = red(fmt.Sprintf("%.2f", transaction.Amount))
			} else {
				amountColored = green(fmt.Sprintf("%.2f", transaction.Amount))
			}
			fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s\t%s\t%s", transaction.Id, transaction.Description, amountColored, transaction.Date, transaction.Category))
		}
	},
}

func init() {
	expenseCmd.AddCommand(expenseListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// add flag to determine if the query is full history
	expenseListCmd.Flags().BoolP("full", "f", false, "List all expenses")
	expenseListCmd.Flags().BoolP("update", "u", false, "Force update of expenses")
}
