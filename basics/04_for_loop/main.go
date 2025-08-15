package main

import "fmt"

// func main() {

// 	var numbers = []int{1, 2, 3, 4, 5, 6}

// 	// Simple iteration over range
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(i)
// 	}

// 	// iteration over collection

// 	for i, value := range numbers {
// 		if value%2 == 0 {
// 			continue
// 		}

// 		fmt.Printf("Odd Number: %v\n", value)

// 		if i == 5 {
// 			break
// 		}
// 	}
// }

// func main() {
// 	rows := 10

// 	for i := 1; i <= rows; i++ {
// 		for j := 1; j <= rows-i; j++ {
// 			fmt.Print(" ")
// 		}

// 		for k := 1; k <= 2*i-1; k++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()

// 	}

// }

// func main() {
// 	for i := range 10 {
// 		fmt.Println(10-i)
// 	}
// }


func main() {
	// usando for como while

	isFinished := false
	var decision int

	for !isFinished {
		fmt.Print("Digite 0 ou 1: ")
		fmt.Scanf("%d", &decision)

		if decision == 1 {
			isFinished = true
		}

	}

	fmt.Print("Você digitou 1, programa finalizado")
}
