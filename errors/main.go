package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Result:", res)
}

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	} else {
		return a / b, nil
	}
}
