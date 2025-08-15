package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().Unix())
	random := rand.New(source)

	randomNumber := random.Intn(20) + 1

	var yourGuess int

	fmt.Println("Guess a Number")

	for {
		fmt.Print("Type a number from 1 to 100 and try to guess: ")
		fmt.Scanf("%d", &yourGuess)

		if randomNumber > yourGuess {
			fmt.Println("Too low guess")
		} else if randomNumber < yourGuess {
			fmt.Println("Too higher guess")
		} else {
			fmt.Printf("That's right, you win, the number is: %v\n",
				randomNumber)
			break
		}
	}
}
