package main

import (
	"fmt"
	 "net/http"
)

func main() {
	fmt.Println("Hello, Go")

	resp, err := http.Get("http://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer resp.Body.Close()

	fmt.Println("HTTP RESPONSE STATUS: ", resp.Status)
}
