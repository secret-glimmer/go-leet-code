package main

import (
	"fmt"
)

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1 // Each child gets at least one candy
	}

	// First pass: left to right
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Second pass: right to left
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	totalCandies := 0
	for _, c := range candies {
		totalCandies += c
	}

	return totalCandies
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ratings1 := []int{1, 0, 2}
	fmt.Println(candy(ratings1)) // Output: 5

	ratings2 := []int{1, 2, 2}
	fmt.Println(candy(ratings2)) // Output: 4
}
