package main

import "fmt"

func main() {
	user := User{
		name:  "Khadga Shrestha",
		email: "khadgadevlove@gmail.com",
	}
	user.log()
	NewPerson("Sakar Subedi", "rakasidebus@gmail.com").log()
}

type User struct {
	name  string
	email string
}

func (user User) log() {
	fmt.Println("name:", user.email, "email:", user.email)
}

func NewPerson(name, email string) User {
	return User{
		name, email,
	}
}
