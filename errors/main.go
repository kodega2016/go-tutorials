package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	res, err := divide(10, 20)
	if err != nil {
		fmt.Println("error occurred:", err)
		os.Exit(1)
	}
	fmt.Println("result:", res)


	result,error:=divide(10,0)
	if error!=nil{
		fmt.Println("error occurred:", error)
		os.Exit(1)
	}
	fmt.Println("result:",result)
}

func divide(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("division by the zero")
	} else {
		return x / y, nil
	}
}