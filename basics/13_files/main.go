package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = getBalanceFromFile()
	var choice int

	if err != nil {
		fmt.Println("An error occurs: ", err)
		return
	}

	fmt.Println("Welcome To Go Bank")

	for {
		menuAndChoice(&choice)

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye")
			return
		}
	}
}

func menuAndChoice(choice *int) {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	fmt.Print("Your Choice: ")

	fmt.Scan(choice)

}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Error to find balance.txt file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Error to parse stored balance to float")
	}

	return balance, nil
}
