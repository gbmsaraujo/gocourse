package main

import (
	"fmt"
	"strconv"
	"strings"
)

func identifyPrefixPostfix(userID, email string) bool {
	return strings.HasPrefix(userID, "UID") && strings.HasSuffix((email), ".io")
}

func containsEducative(email string) bool {
	return strings.Contains(email, "educative")
}

func maskUserName(email string) string {
	splitUserName := strings.Split(email, "@")
	userName := []rune(splitUserName[0])

	emailDomain := splitUserName[1]
	asteristicWord := strings.Repeat("*", len(userName)-2)
	response := string(userName[0]) + asteristicWord + string(userName[len(userName)-1]) + "@" + emailDomain
	return response
}

func indexOfAtSymbol(email string) int {
	// Implement this function
	indexEmail := strings.Index(email, "@")
	return indexEmail
}

func trimAndSplitUserID(userID string) string {
	userIDSplitted := strings.Split(strings.TrimSpace(userID), "-")
	return userIDSplitted[1]
}

func convertStringToInt(str string) int {
	value, _ := strconv.Atoi(str)

	return value

}

func main() {
	// Test your functions here
	fmt.Println(identifyPrefixPostfix("UID-0123", "evangeline@educative.io")) // true
	fmt.Println(containsEducative("evangeline@educative.io"))                 // true
	fmt.Println(maskUserName("evangeline@educative.io"))                      // e******e@educative.io
	fmt.Println(indexOfAtSymbol("evangeline@educative.io"))                   // 10
	fmt.Println(trimAndSplitUserID("UID-0123"))                               // 0123
	fmt.Println(convertStringToInt("123"))                                    // 123
}
