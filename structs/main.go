package main

import "fmt"

func main() {
	user := User{
		id:   UserID("abc"),
		name: "Khadga Bahdur Shrestha",
		role: "Software Developer",
		age:  27,
		address: Address{
			country: "Nepal",
			street:  "Mikaljung morang",
			name:    "Mikaljung morang,madhumalla",
		},
	}

	fmt.Println("user:", user)
	fmt.Println("name:", user.name)
	fmt.Println("age:", user.age)
	fmt.Println("address:", user.address.name)
}

type UserID string

type User struct {
	id      UserID
	name    string
	role    string
	age     int
	address Address
}

type Address struct {
	name    string
	country string
	street  string
}
