package main

import "fmt"

func main() {
	setup()
	total := getSum(10, 20)
	fmt.Println("the sum of 10 and 20 is", total)

	sum, sub := getSumAndSub(20, 30)
	fmt.Println("sum:", sum, "sub:", sub)

	res := getSquare(20)
	fmt.Println("result of the sqaure:", res)

	performAction("Explore go for software development", logMessage)
}

func init() {
	fmt.Println("starting the program...")
}

func setup() {
	fmt.Println("setting up server...")
}

func getSum(a, b int) int {
	return a + b
}

func getSumAndSub(a, b int) (int, int) {
	return a + b, a - b
}

func getSquare(num int) (sqaure int) {
	sqaure = num * num
	return
}

func performAction(params string, action func(string)) {
	action(params)
}

func logMessage(msg string) {
	fmt.Println("log:", msg)
}
