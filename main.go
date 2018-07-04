package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	initLogs()
	initConfig()

	router := mux.NewRouter()
	initRouter(router)

	http.ListenAndServe(":8080", router)
}

func initRouter(router *mux.Router) {
	router.HandleFunc("/v1/balance", getBalance)
	router.HandleFunc("/v1/transactions", getTransactions)
	router.HandleFunc("/v1/bills", getBills)
	router.HandleFunc("/v1/summary", getSummary)
}
