package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var (
	baseUrl         = "https://api.starlingbank.com/api/"
	balanceEndpoint = "v1/accounts/balance"

	httpClient = &http.Client{
		Timeout: time.Second * 10,
	}
)

type (
	starlingAccount struct {
		ClearedBalance      float64 `json:"clearedBalance"`
		EffectiveBalance    float64 `json:"effectiveBalance"`
		PendingTransactions float64 `json:"pendingTransactions"`
		AvailableToSpend    float64 `json:"availableToSpend"`
		AcceptedOverdraft   float64 `json:"acceptedOverdraft"`
		Currency            string  `json:"currency"`
		Amount              float64 `json:"amount"`
	}
)

func updateBalance() starlingAccount {
	Info.Println("Updating")
	start := time.Now()

	var body starlingAccount

	req, err := http.NewRequest(http.MethodGet, baseUrl+balanceEndpoint, nil)

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

	return body
}
