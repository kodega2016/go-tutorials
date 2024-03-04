package main

import "fmt"

func main() {
	u1 := User{
		name:  "Sakar Subedi",
		email: "sakar@email.com",
		score: 100,
	}

	u2 := User{
		name:  "Sabin Nepal",
		email: "sabin@email.com",
		score: 105,
	}

	fmt.Println(u1, u2)
}

type User struct {
	name  string
	email string
	score int
}

func (u User) String() string {
	return fmt.Sprintf("Name:%s,Email:%s, Score:%d\n", u.name, u.email, u.score)
}
