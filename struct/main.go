package main

import (
	"fmt"
	// "unsafe"
	"structs_app/user"
)

func main() {
	userFirstName := getUserData("Please enter your first Name: ")
	userLastName := getUserData("Please enter your last Name: ")
	userBirthdate := getUserData("Please enter your birthdate (dd/mm/yyyy): ")

	appUser, error := user.New(userFirstName, userLastName, userBirthdate)

	if error != nil {
		fmt.Println(error)
		return
	}

	appAdmin := user.NewAdmin("test@example", "pdw2324")

	appAdmin.UserDetails()

	appUser.UserDetails()
	appUser.ClearUserName()
	appUser.UserDetails()

}

func getUserData(msg string) string {
	var data string
	fmt.Print(msg)
	fmt.Scanln(&data)
	return data
}
