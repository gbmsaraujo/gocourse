package main

import (
	"errors"
	"fmt"
)

func main() {
	var selectedMonth int
	fmt.Print("Choose a month: ")
	fmt.Scan(&selectedMonth)
	season, err := seasons(selectedMonth)

	if err !=nil {
		fmt.Println(err)
		return
	}

	fmt.Println(season)
}

func seasons(month int) (string, error) {
	switch month {
	case 1, 2, 12:
		return "Winter", nil
	case 3, 4, 5:
		return "Spring", nil
	case 6, 7, 8:
		return "Summer", nil
	case 9, 10, 11:
		return "Autumn", nil
	default:
		return  "", errors.New("Invalid month: type from 1 to 12, where 1 is january and 12 is december")
	}
}

