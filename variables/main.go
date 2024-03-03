package main

import "fmt"

func main() {
	// explicit declaration
	var name string = "Software Engineering with Golang"
	fmt.Println("Book name:", name)
	// implicit declaration
	age := 15
	// age="nice"
	fmt.Println("age:", age)

	username, role, status := "kodega2016", "software developer", "active"
	fmt.Println("username:", username, "role:", role, "status:", status)

	var (
		company_name    string = "PortPro Nepal"
		company_size    int    = 24
		company_address string = "Sanepa"
	)
	fmt.Println("company name:", company_name, "company size:", company_size, "company address:", company_address)

	var book_name string
	var total_marks int
	var is_active bool
	var score float32

	fmt.Println("book_name:", book_name, "total marks:", total_marks, "is_active:", is_active, "score:", score)
}
