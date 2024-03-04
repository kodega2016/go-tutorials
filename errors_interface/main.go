package main

import (
	"fmt"
	"os"
)

func main() {
	a := Auth{
		email:    "user@user.com",
		password: "pass123",
	}

	is_logged_in, err := a.login()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if is_logged_in {
		fmt.Println("you are successfully logged in")
	}

}

type Auth struct {
	email    string
	password string
}

func (a Auth) login() (bool, error) {
	if a.email == "admin@admin.com" && a.password == "password" {
		return true, nil
	}
	return false, a
}

func (a Auth) Error() string {
	return fmt.Sprintf("failed to login with email %s and password:%s", a.email, a.password)
}
