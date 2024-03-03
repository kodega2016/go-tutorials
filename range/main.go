package main

import "fmt"

func main() {
	var scores []int = []int{10, 20, 45, 54, 65}
	totalScores := 0
	for _, score := range scores {
		totalScores += score
	}

	fmt.Println("total scores:", totalScores)
	fmt.Println("average score:", totalScores/len(scores))

	var userDoc map[string]any = map[string]any{
		"name":          "khadga bahadur shrestha",
		"score":         12,
		"is_active":     true,
		"leave_balance": 2.4,
	}

	fmt.Println(userDoc)
	for key, value := range userDoc {
		fmt.Println(key, value)
	}
}
