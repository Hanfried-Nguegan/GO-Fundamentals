package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u *User) outputPlayerDetails() {
	fmt.Printf(" Firstname: %s\n Lastname: %s\n Birthdate: %s\n CreatedAt %s\n", u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) clearUsername() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthdate string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *User

	appUser = newUser(userFirstName, userLastName, userBirthdate)

	appUser.outputPlayerDetails()
	appUser.clearUsername()
	appUser.outputPlayerDetails()
	// ... do something awesome with that gathered data!
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
