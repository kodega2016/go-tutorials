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
}
