package main

type (
	Summary struct {
		Income   float64
		Expences float64
		Savings  float64
	}
)

func calculateSummary() Summary {
	var summary Summary

	summary.Income = 2480
	summary.Expences = calculateBills().Amount + proccessStarlingTransactions(strarlingTransactions()).Amount

	summary.Savings = summary.Income + summary.Expences

	return summary
}
