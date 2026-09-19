package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u *User) OutputPlayerDetails() {
	fmt.Printf(" Firstname: %s\n Lastname: %s\n Birthdate: %s\n CreatedAt %s\n", u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}

func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("Firstname, Lastname and Birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
