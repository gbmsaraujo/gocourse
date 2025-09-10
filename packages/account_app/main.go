package main

import (
	"account_app/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	var choice int

	if err != nil {
		fmt.Println("An error occurs: ", err)
		return
	}

	fmt.Println("Welcome To Go Bank")
	fmt.Println(randomdata.FirstName(2))

	for {
		presentOptions()
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is ", accountBalance)
		case 2:
			fmt.Println("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be Greater than 0.")
				continue
			}
			accountBalance += depositAmount

			fmt.Println("Your balance updated is ", accountBalance)
			fileops.WriteValueToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Println("withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient founds")
				continue
			}
			accountBalance -= withdrawAmount

			fmt.Println("Your balance updated is ", accountBalance)
			fileops.WriteValueToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye")
			return
		}
	}
}
