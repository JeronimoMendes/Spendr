package tracker

import (
	"fmt"
	"strings"
	"time"

	"bytes"
	"encoding/gob"
	"os"

	"github.com/JeronimoMendes/expense-track/pkg/gc_client"
)

type Expense struct {
	Id string
	Amount float64
	Description string
	Currency string
	Date time.Time
	Category string
}

type ExpenseTracker struct {
	GCClient *gc_client.GoCardlessClient
	Expenses []Expense
	Categories []string
	LastExpensesUpdate time.Time
}

func NewExpenseTracker(gcClient *gc_client.GoCardlessClient) *ExpenseTracker {
	data, err := os.ReadFile(getStateFilePath())
	if err != nil {
		fmt.Println("No state file found, creating a new one.")
		return &ExpenseTracker{
			GCClient: gcClient,
			Expenses: []Expense{},
			Categories: []string{},
			LastExpensesUpdate: time.Time{},
		}
	}

	var expenseTracker ExpenseTracker
	dec := gob.NewDecoder(bytes.NewBuffer(data))
	if err := dec.Decode(&expenseTracker); err != nil {
		panic(err)
	}

	expenseTracker.GCClient = gcClient
	return &expenseTracker
}

func (tracker *ExpenseTracker) GetExpenses(accountID string, full bool) []Expense {
	if time.Since(tracker.LastExpensesUpdate).Hours() > 24 {
		tracker.updateExpenses(accountID)
	}
	if full {
		return tracker.Expenses
	}

	var expenses []Expense
	for _, expense := range tracker.Expenses {
		if expense.Date.Month() == time.Now().Month() {
			expenses = append(expenses, expense)
		}
	}

	return expenses
}

func (tracker *ExpenseTracker) GetCategories() []string {
	return tracker.Categories
}

func (tracker *ExpenseTracker) CreateCategory(category string) {
	tracker.Categories = append(tracker.Categories, category)
	tracker.Save()
}

func (tracker *ExpenseTracker) DeleteCategory(category string) {
	for i, cat := range tracker.Categories {
		if cat == category {
			tracker.Categories = append(tracker.Categories[:i], tracker.Categories[i+1:]...)
		}
	}
	tracker.Save()
}

func (tracker *ExpenseTracker) Save() {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(tracker); err != nil {
		panic(err)
	}

	if err := os.WriteFile(getStateFilePath(), buf.Bytes(), 0755); err != nil {
		panic(err)
	}
}

func getStateFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return home + "/.expense_track/data.gob"
}

func (tracker *ExpenseTracker) updateExpenses(accountID string) {
	gcExpenses := tracker.GCClient.GetExpenses(accountID)
	var updated bool

	for _, gcExpense := range gcExpenses {
		if tracker.expenseExists(gcExpense.Id) == -1 {
			updated = true
			expense := Expense{
				Id: gcExpense.Id,
				Amount: gcExpense.Amount,
				Description: gcExpense.Description,
				Currency: gcExpense.Currency,
				Date: gcExpense.Date,
				Category: "",
			}
			tracker.Expenses = append(tracker.Expenses, expense)
		}
	}

	tracker.LastExpensesUpdate = time.Now()
	if updated {
		tracker.Save()
	}
}

func (tracker *ExpenseTracker) ResetExpenses() {
	tracker.Expenses = []Expense{}
	tracker.LastExpensesUpdate = time.Time{}
	tracker.Save()
}

func (tracker *ExpenseTracker) expenseExists(id string) int {
	found := false
	expenseIndex := -1
	for i, expense := range tracker.Expenses {
		if strings.HasPrefix(expense.Id, id) {
			if found {
				panic("Multiple expenses found with that ID prefix")
			}

			found = true
			expenseIndex = i
			if expense.Id == id {
				break
			}
		}
	}
	return expenseIndex
}

func (tracker *ExpenseTracker) categoryExists(category string) bool {
	for _, cat := range tracker.Categories {
		if cat == category {
			return true
		}
	}
	return false
}


func (tracker *ExpenseTracker) CategoriseExpense(id string, category string) {
	expenseIndex := tracker.expenseExists(id)
	if expenseIndex == -1 {
		panic("Expense does not exist")
	}

	if !tracker.categoryExists(category) {
		panic("Category does not exist")
	}

	tracker.Expenses[expenseIndex].Category = category
	tracker.Save()
}

func (tracker *ExpenseTracker) GetExpensesByCategory(category string) []Expense {
	var expenses []Expense
	for _, expense := range tracker.Expenses {
		if expense.Date.Month() == time.Now().Month() && expense.Category == category {
			expenses = append(expenses, expense)
		}
	}
	return expenses
}

func (tracker *ExpenseTracker) GetTotalExpensesByCategory(category string) float64 {
	var total float64
	for _, expense := range tracker.Expenses {
		if expense.Date.Month() == time.Now().Month() && expense.Category == category {
			total += expense.Amount
		}
	}
	return total
}
