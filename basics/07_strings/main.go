package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"

	fmt.Printf("T/F? \nDoes the string \"%s\" have prefix %s? \n", str, "Th")
	// fmt.Printf("\n%t\n\n", strings.HasPrefix(str, "Th"))    // Finding prefix
	// fmt.Printf("\n%t\n\n", strings.HasSuffix(str, "ing"))

	sentence := "I love you"

	fmt.Printf("The sentence contains you? %v \n", strings.Contains(sentence, "you"))

	fmt.Printf("Index of you? %v \n", strings.Index(sentence, "you"))

	sentenceTwo:= "I love you so much, only you"

	fmt.Printf("Last Index of you? %v \n", strings.LastIndex(sentenceTwo, "you"))

	var strTwo string = "Hello, how is it going, Hugo?"
    var manyG = "gggggggggg"
    fmt.Printf("Number of H's in %s is: ", strTwo)
    fmt.Printf("%d\n", strings.Count(strTwo, "H"))         // count occurences 
    fmt.Printf("Number of double g's in %s is: ", manyG)
    fmt.Printf("%d\n", strings.Count(manyG, "gg"))      // count occurences
}
