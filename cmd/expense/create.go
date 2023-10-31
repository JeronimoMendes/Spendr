package expense

import (
	"fmt"
	"strconv"
	"time"

	"github.com/JeronimoMendes/spendr/pkg/gc_client"
	"github.com/JeronimoMendes/spendr/pkg/tracker"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new expense",

	Run: func(cmd *cobra.Command, args []string) {
		secretID := viper.GetString("gc_secret_id")
		secretKey := viper.GetString("gc_secret_key")

		gc := gc_client.NewClient(secretKey, secretID)
		tracker := tracker.NewExpenseTracker(gc)

		fmt.Println("Creating new expense...")

		prompt := promptui.Prompt{
			Label: "Description",
		}

		description, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		prompt = promptui.Prompt{
			Label: "Amount",
		}

		amountInput, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		amount, err := strconv.ParseFloat(amountInput, 64)
		if err != nil {
			fmt.Printf("%s is not a valid float\n", err)
			return
		}

		choose := promptui.Select{
			Label: "Select currency",
			Items: []string{"EUR", "USD", "GBP"},
		}

		_, currency, err := choose.Run()	
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		prompt = promptui.Prompt{
			Label: "Date",
		}

		dateInput, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		// convert dateInput to time.Time
		date, err := time.Parse("2006-01-02", dateInput)
		if err != nil {
			fmt.Printf("%s is not a valid date\n", err)
			return
		}

		categories := []string{}
		for _, category := range tracker.Categories {
			categories = append(categories, category.Name)
		}
		choose = promptui.Select{
			Label: "Select Category",
			Items: categories,
		}

		_, category, err := choose.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		confirmString := fmt.Sprintf("Create expense with description %s, amount %s, date %s and category %s?", description, amountInput, dateInput, category)

		prompt = promptui.Prompt{
			Label:     confirmString,
			IsConfirm: true,
		}

		confirm, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if confirm == "y" {
			tracker.CreateExpense(description, amount, currency, date, category)
			fmt.Println("Expense created.")
		} else {
			fmt.Println("Expense not created.")
		}
	},
}

func init() {
	expenseCmd.AddCommand(createCmd)
}