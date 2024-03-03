package main

import "fmt"

func main() {
	var auth BaseAuth
	auth = Auth{}
	is_logged_in := auth.login("khadgadevlove@gmail.com", "password")
	if is_logged_in {
		fmt.Println("you are logged in successfully...")
	}
}

type BaseAuth interface {
	login(email, password string) bool
}
type Auth struct{}

// login implements BaseAuth.
func (Auth) login(email string, password string) bool {
	fmt.Println("login with ", email, password)
	return true
}
