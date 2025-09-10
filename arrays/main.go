package main

import "fmt"

// func main() {
// 	// 1. Nosso array original, fixo, com 6 elementos.
// 	// Em memória: [10, 20, 30, 40, 50, 60]
// 	meuArray := [6]int{10, 20, 30, 40, 50, 60}
// 	fmt.Printf("Array Original: len=%d, cap=%d, %v\n", len(meuArray), cap(meuArray), meuArray)

// 	// 2. Criamos um slice que é uma "visão completa" do array.
// 	sliceCompleto := meuArray[:] // é o mesmo que meuArray[0:6]
// 	// Ponteiro aponta para o índice 0.
// 	// len = 6 - 0 = 6
// 	// cap = 6 - 0 = 6
// 	fmt.Printf("Slice Completo: len=%d, cap=%d, %v\n", len(sliceCompleto), cap(sliceCompleto), sliceCompleto)

// 	// 3. Criamos uma "visão" do meio do array.
// 	visaoDoMeio := meuArray[2:4] // Pega os elementos de índice 2 e 3.
// 	// Ponteiro aponta para o índice 2 (valor 30).
// 	// len = 4 - 2 = 2
// 	// cap = 6 - 2 = 4 (porque do índice 2 até o fim do array, cabem 4 elementos)
// 	fmt.Printf("Visão do Meio: len=%d, cap=%d, %v\n", len(visaoDoMeio), cap(visaoDoMeio), visaoDoMeio)
// }

func main() {
	prices := []float64{}
	prices = append(prices, 5.99)
	prices = append(prices, 10.95)
	prices = append(prices, 20.10)
	prices = append(prices, 32.99)
	fmt.Println(prices)
}
