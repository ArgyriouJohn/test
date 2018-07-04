package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var (
	transactionsEndpoint = "v1/transactions"
)

type (
	transactionsSummary struct {
		Amount       float64       `json:"amount"`
		Transactions []Transaction `json:"transactions"`
	}

	Root struct {
		Embedded struct {
			Transactions []Transaction `json:"transactions"`
		} `json:"_embedded"`
	}

	Transaction struct {
		ID        string    `json:"id"`
		Currency  string    `json:"currency"`
		Amount    float64   `json:"amount"`
		Direction string    `json:"direction"`
		Created   time.Time `json:"created"`
		Narrative string    `json:"narrative"`
		Source    string    `json:"source"`
		Balance   float64   `json:"balance"`
	}
)

func strarlingTransactions() []Transaction {
	Info.Println("Updating")
	start := time.Now()

	var body Root

	req, err := http.NewRequest(http.MethodGet, baseUrl+transactionsEndpoint, nil)

	if err != nil {
		Error.Println("Request, Error : ", err)
	} else {
		req.Header.Set("Authorization", "Bearer "+conf.Starling.StarlingToken)
		res, err := httpClient.Do(req)

		if err != nil {
			Error.Println("Response, Error : ", err)
		} else {
			defer res.Body.Close()

			if res.StatusCode != 200 {
				Error.Println("Http Response Error", res.StatusCode)
			} else {

				err = json.NewDecoder(res.Body).Decode(&body)
				if err != nil {
					Error.Println("Decode, Error : ", err)
				}
			}
		}
	}

	stop := time.Now()
	Info.Println("Finished Updating in : ", stop.Unix()-start.Unix(), " seconds")

	return body.Embedded.Transactions
}

func proccessStarlingTransactions(Transactions []Transaction) transactionsSummary {
	var ts transactionsSummary
	amount := 0.0

	//Remove Outbound
	for i, transaction := range Transactions {
		if transaction.Direction == "INBOUND" {
			Transactions = append(Transactions[:i], Transactions[i+1:]...)
		} else {
			if transaction.Narrative != "Revolut" {
				amount = amount + transaction.Amount
			}
		}
	}

	ts.Amount = amount
	ts.Transactions = Transactions

	return ts
}
