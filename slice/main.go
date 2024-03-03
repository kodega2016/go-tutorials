package main

import "fmt"

func main() {
	var users []string
	users = []string{
		"sakar", "sabin", "asok", "himal",
	}
	fmt.Println("users:", users)

	fmt.Println("total users:", len(users))
	fmt.Println("last user:", users[len(users)-1])

	// append to the slice
	users = append(users, "khadga", "bharat")
	fmt.Println("users:", users)
}
