package main

import (
	"fmt"
	"os"
)

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go bank !")

	for {
		fmt.Println("What do you want to do ?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				//return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated ! New Amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Your Withdrawal: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. must be greater than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("INvalid amount. You can't withdraw more than what you have")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("You withdrew:", withdrawAmount, "Your new balance is:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Program Exited")
			fmt.Println("Thanks for choosing our bank")
			return
			//break
		}
	}
}
