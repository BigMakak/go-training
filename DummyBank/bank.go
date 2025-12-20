package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const balanceFileName = "balance.txt"

func main() {
	var accountBalance, err = fileops.ReadFloatFromFile(balanceFileName, 0.0)
	if err != nil {
		fmt.Println("ERROR: Starting with a balance of $0.0")
		fmt.Println(err)
		fmt.Println("--------------------------------")
	}

	fmt.Println("Welcome to the Big Go Bank!")
	fmt.Println("You can call for assistance at: " + randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:

			fmt.Println("Your balance is: $", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be greater than zero.")
				continue
			}

			accountBalance += depositAmount
			err := fileops.WriteFloatToFile(balanceFileName, accountBalance)
			if err != nil {
				fmt.Println("Error updating balance file after deposit.")
				return
			}
			fmt.Println("Deposit successful! New balance is: $", accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Withdrawal amount must be greater than zero.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds! Your current balance is: $", accountBalance)
				continue
			}

			accountBalance -= withdrawAmount
			err := fileops.WriteFloatToFile(balanceFileName, accountBalance)
			if err != nil {
				fmt.Println("Error updating balance file after deposit.")
				return
			}
			fmt.Println("Withdraw successful! New balance is: $", accountBalance)
		default:
			fileops.WriteFloatToFile(balanceFileName, accountBalance)
			fmt.Println("Thank you for banking with us. Goodbye!")
			fmt.Println("Exiting Bank...")
			return
		}
		fmt.Println("Your choice is:", choice)
	}
}
