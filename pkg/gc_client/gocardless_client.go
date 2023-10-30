package gc_client

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var GC_API_URL = "https://ob.gocardless.com/api/v2"

type GoCardlessClient struct {
	AccessToken string
}

type ExpenseGC struct {
	Id string
	Amount float64
	Description string
	Currency string
	Date time.Time
}

type NewTokenResponse struct {
	Access string
}

type ListTransactionsResponse struct {
	Transactions TransactionResponse
}

type TransactionResponse struct {
	Booked []Transaction
	Pending []Transaction
}

type Transaction struct {
	InternalTransactionId string
	ValueDate string
	RemittanceInformationUnstructured string
	TransactionAmount TransactionAmountResponse
}

type TransactionAmountResponse struct {
	Amount string
	Currency string
}

func NewClient(SecretKey string, SecretID string) *GoCardlessClient {
	res, error := http.PostForm(GC_API_URL + "/token/new/", url.Values{
		"secret_id": {SecretID},
		"secret_key": {SecretKey},
	})

	if error != nil {
		panic(error)
	}

	if res.StatusCode != 200 {
		fmt.Println("Error getting new token")
		fmt.Println(res)
		panic(res)
	}

	var data NewTokenResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Error decoding response in new token")
		panic(err)
    }

	return &GoCardlessClient{AccessToken: data.Access}
}

func (c *GoCardlessClient) GetExpenses(account string) []ExpenseGC {
	// res, error := http.Get(fmt.Sprintf("%s/accounts/%s/transactions/", GC_API_URL, account))
	httpClient := &http.Client{}

	req, error := http.NewRequest("GET", fmt.Sprintf("%s/accounts/%s/transactions/", GC_API_URL, account), nil)
	if error != nil {
		fmt.Println("Error creating new request")
		panic(error)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	res, error := httpClient.Do(req)
	if error != nil {
		fmt.Println("Doing request for list transactions")
		panic(error)
	}

	var data ListTransactionsResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Error decoding response in list transactions")
		panic(err)
	}

	// create a list of ExpenseGC
	var expenses []ExpenseGC
	for _, transaction := range data.Transactions.Booked {
		dateAsTime, err := time.Parse("2006-01-02", transaction.ValueDate)
		if err != nil {
			fmt.Println("Error parsing date")
			panic(err)
		}
		amountAsInt, err := strconv.ParseFloat(transaction.TransactionAmount.Amount, 64)

		expenses = append(expenses, ExpenseGC{
			Id: transaction.InternalTransactionId,
			Amount: amountAsInt,
			Description: transaction.RemittanceInformationUnstructured,
			Currency: transaction.TransactionAmount.Currency,
			Date: dateAsTime,
		})
	}

	return expenses
}