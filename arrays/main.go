package main

import "fmt"

func main() {
	var users [10]string
	fmt.Println("users:", users)
	users = [10]string{"sakar", "samir", "sabin"}
	fmt.Println("users:", users)

	fmt.Println("total users:", len(users))

	for i := 0; i < len(users); i++ {
		fmt.Println(i, users[i])
	}
}
