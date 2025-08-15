package main

import "fmt"

func main() {
// LABEL1: // adding a label for location
	for i := 0; i <= 5; i++ { // outer loop
		for j := 0; j <= 5; j++ { // inner loop
			if j == 4 {
				fmt.Printf("==> j é %v. continue LABEL1!\n", j)
				break // jump to the label
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
		fmt.Println("Fim do laço interno para i =", i)
	}

}
