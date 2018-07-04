package main

import (
	"time"
	"net/http"
	"encoding/json"
)

var (
	monzoBaseUrl = "https://api.monzo.com"
	monzoTransactionsUrl = "/transactions?account_id="+conf.Monzo.CurrentAccount.CurrentAccount
)

func monzoTransactions(){
	Info.Println("Updating")
	start := time.Now()

	var body Root

	req, err := http.NewRequest(http.MethodGet, monzoBaseUrl+monzoTransactionsUrl, nil)

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