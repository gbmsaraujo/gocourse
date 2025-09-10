package main

import "fmt"

func main() {
	var age int

	fmt.Println("Type your age: ")

	fmt.Scan(&age)

	editAdultYears(&age)

	if age >= 0 {
		fmt.Println("You're adult and can access the bar")
	} else {
		fmt.Println("You're not adult and can't drink")
	}

}

func editAdultYears(age *int) {
	*age = *age - 18
}
