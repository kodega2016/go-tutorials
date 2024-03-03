package main

import "fmt"

func main() {
	var address *int
	age := 10
	address = &age
	fmt.Println("age:", *address, "address:", address)

	c := Counter{}
	c.increment()
	c.logCounter()
}

type Counter struct {
	count int
}

func (c *Counter) increment() {
	c.count++
}

func (c Counter) logCounter() {
	fmt.Println("The value of count is", c.count)
}
