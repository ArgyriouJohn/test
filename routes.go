package main

import (
	"encoding/json"
	"net/http"
)

type (
	response struct {
		Status   string      `json:"status"`
		Data     interface{} `json:"data,omitempty"`
		Message  string      `json:"message,omitempty"`
		Metadata metadata    `json:"metadata"`
	}

	metadata struct {
		CreatedAt  string  `json:"createdAt"`
		TotalItems *int    `json:"totalItems,omitempty"`
		Sort       *string `json:"sort,omitempty"`
	}
)

func getBalance(w http.ResponseWriter, r *http.Request) {
	balance := updateBalance()

	json.NewEncoder(w).Encode(response{
		Status:  "success",
		Data:    balance,
		Message: "",
		Metadata: metadata{
			now(), nil, nil},
	})
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	transactionsSummary := proccessStarlingTransactions(strarlingTransactions())

	json.NewEncoder(w).Encode(response{
		Status:  "success",
		Data:    transactionsSummary,
		Message: "",
		Metadata: metadata{
			now(), nil, nil},
	})
}

func getBills(w http.ResponseWriter, r *http.Request) {
	billsSummary := calculateBills()

	json.NewEncoder(w).Encode(response{
		Status:  "success",
		Data:    billsSummary,
		Message: "",
		Metadata: metadata{
			now(), nil, nil},
	})
}

func getSummary(w http.ResponseWriter, r *http.Request) {
	summary := calculateSummary()

	json.NewEncoder(w).Encode(response{
		Status:  "success",
		Data:    summary,
		Message: "",
		Metadata: metadata{
			now(), nil, nil},
	})
}
