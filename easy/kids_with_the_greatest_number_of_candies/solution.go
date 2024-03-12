package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	for _, kidCandies := range candies {
		if maxCandies < kidCandies {
			maxCandies = kidCandies
		}
	}

	flags := []bool{}
	for _, kidCandies := range candies {
		flags = append(flags, kidCandies+extraCandies >= maxCandies)
	}

	return flags
}

func main() {
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))
}
