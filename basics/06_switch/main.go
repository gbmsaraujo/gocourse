package main

import (
	"fmt"
)

func main() {
	fruit := "pineapple"

	switch fruit {
	case "apple":
		printMessage("It's an apple")
	case "banana":
		printMessage("It's a banana")
	default:
		printMessage("It's an unknown fruit")
	}

	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		printMessage("It's a weekday.")

	case "Saturday", "Sunday":
		printMessage("It's a weekend.")
	default:
		printMessage("Invalid day")
	}

	grade := 8.0

	switch {
	case grade >= 9:
		printMessage("Perfect grade")
	case grade >= 8 && grade < 9:
		printMessage("Great grade")
	case grade >= 7 && grade < 8:
		printMessage("Good grade")
	case grade >= 6 && grade < 7:
		printMessage("Regular grade")
	default:
		printMessage("Failed!")

	}
	var anyList = []int{1,2,3}
	
	checkType(2)
	checkType(2.3)
	checkType("OI")
	checkType(false)
	checkType(anyList)
	fmt.Printf("é do tipo: %T\n", anyList)

	anyList = append(anyList, 4)

}

func printMessage(message string) {
	fmt.Println(message)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		printMessage("It's an integer")
	case float64:
		printMessage("It's a float")
	case string:
		printMessage("It's a string")
	case bool:
		printMessage("It's a bool")
	default:
		printMessage("Unknown Type")
	}
}
