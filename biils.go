package main

import "time"

const (
	Rent     = "RENT"
	Tax      = "TAX"
	Utility  = "UTILITY"
	Personal = "PERSONAL"
)

type (
	BillsSummary struct {
		Amount float64
		Bills  []Bill
	}

	Bill struct {
		Description string
		Type        string
		Date        time.Time
		Amount      float64
	}
)

func calculateBills() BillsSummary {
	var billsSummary BillsSummary
	var bills []Bill
	var amount float64

	rent := Bill{
		Description: "Rent",
		Type:        Rent,
		Date:        formatTime("2018-06-02T00:00:00.000Z"),
		Amount:      -650,
	}

	councilTax := Bill{
		Description: "Council Tax",
		Type:        Tax,
		Date:        formatTime("2018-06-02T00:00:00.000Z"),
		Amount:      -150,
	}

	electricity := Bill{
		Description: "Electricity",
		Type:        Utility,
		Date:        formatTime("2018-06-02T00:00:00.000Z"),
		Amount:      -20.99,
	}

	water := Bill{
		Description: "Water",
		Type:        Utility,
		Date:        formatTime("2018-06-02T00:00:00.000Z"),
		Amount:      -19.2,
	}

	mobile := Bill{
		Description: "Mobile",
		Type:        Utility,
		Date:        formatTime("2018-06-02T00:00:00.000Z"),
		Amount:      -11,
	}

	gym := Bill{
		Description: "Gym",
		Type:        Personal,
		Date:        formatTime("2018-06-02T00:00:00.000Z"),
		Amount:      -20.99,
	}

	bills = append(bills,
		rent,
		councilTax,
		electricity,
		water,
		mobile,
		gym,
	)

	billsSummary.Bills = bills

	for _, bill := range bills {
		amount = amount + bill.Amount
	}

	billsSummary.Amount = amount

	return billsSummary
}
