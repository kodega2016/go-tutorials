package main

import "fmt"

func main() {
	age := 10

	if age < 18 {
		fmt.Println("you are not illigible to this site")
	} else {
		fmt.Println("you are allowed to this site")
	}

	role := "admin"
	switch role {
	case "developer":
		fmt.Println("you are a developer")
	case "admin":
		fmt.Println("you are a admin")
		fallthrough
	case "manager":
		fmt.Println("you are a manager")
	default:
		fmt.Println("you are a guest")
	}

	if num := 8; num < 0 {
		fmt.Println("number is less than 0")
	} else if num < 10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println("number has different values")
	}
}
