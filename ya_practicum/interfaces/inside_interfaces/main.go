package main

import (
	"fmt"
	"time"
)

type Stringer interface {
	String() string
}

type User struct {
	Email      string
	Password   string
	LastAccess time.Time
}

func (u User) String() string {
	return "user with email: " + u.Email
}

func Printf(v Stringer) {
	fmt.Print("It's the type implemented Stringer, " + v.String())
}

func main() {
	u := User{Email: "email@example.com"}
	Printf(u)
}
