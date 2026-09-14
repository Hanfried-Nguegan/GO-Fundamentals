package main

import "fmt"

func main() {
	var username string
	username = "Hanfried Nguegan"

	var isAdmin bool
	isAdmin = true

	var permissions int
	permissions = 0x1F

	var costPerSMS float64
	costPerSMS = 0.05

	fmt.Println("Username:", username)
	fmt.Println("isAdmin:", isAdmin)
	fmt.Println("Permissions:", permissions)
	fmt.Println("CostPerSMS:", costPerSMS)
}
