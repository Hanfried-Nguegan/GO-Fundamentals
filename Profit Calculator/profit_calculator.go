package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue := getUserInput("Revenue: ")
	// fmt.Print("Enter Revenue: ")
	// fmt.Scan(&revenue)

	expenses := getUserInput("Expenses: ")
	// fmt.Print("Enter Expenses: ")
	// fmt.Scan(&expenses)

	taxRate := getUserInput("Tax Rate: ")
	// fmt.Print("Enter Tax Rate: ")
	// fmt.Scan(&taxRate)

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := (ebt) * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}
