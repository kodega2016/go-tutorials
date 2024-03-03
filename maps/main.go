package main

import "fmt"

func main() {
	var userDoc map[string]string
	userDoc = make(map[string]string)
	userDoc["name"] = "Khadga Bahadur shrestha"
	userDoc["role"] = "Software Developer"
	userDoc["address"] = "Madhumalla morang"
	fmt.Println(userDoc)

	// delete(userDoc, "name")
	userDoc["has_data"] = "true"

	name, hasName := userDoc["name"]
	if hasName {
		fmt.Println("name:", name)
	} else {
		fmt.Println("name is not set")
	}

	for key, value := range userDoc {
		fmt.Println(key, value)
	}
}
