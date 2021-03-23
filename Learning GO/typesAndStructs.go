package main

import (
	"log"
	"time"
)

//! global scope
var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Sara",
		LastName:  "Peña",
	}

	log.Println(user.FirstName, user.LastName)

}

//! Like JAVA: Public, Private
// ! GO: "User" is a public func, thus visible outside Main Package. "user" is a private func, thus only visible inside the package.
