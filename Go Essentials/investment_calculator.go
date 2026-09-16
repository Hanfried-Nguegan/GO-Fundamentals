package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationrate = 6.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate = 2.5

	fmt.Print("Enter Investment Amount")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationrate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)

	formattedRV := fmt.Sprintf("Future Real Value: %.1f\n", futureRealValue)

	// fmt.Println("Future value:", futureValue)
	//fmt.Printf("Future Value: %.1f\nFuture Real Value: %.1f", futureValue, futureRealValue)
	// fmt.Println("Future Real Value (adjusted for Inflation)", futureRealValue)

	fmt.Print(formattedFV, formattedRV)
}
