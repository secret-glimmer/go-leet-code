package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	n := len(digits)

	// Start from the last digit and work backwards
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	// If all digits were 9, we need to add a new digit at the front
	return append([]int{1}, digits...)
}

func main() {
	// Test cases
	fmt.Println(plusOne([]int{1, 2, 3}))    // Output: [1, 2, 4]
	fmt.Println(plusOne([]int{4, 3, 2, 1})) // Output: [4, 3, 2, 2]
	fmt.Println(plusOne([]int{9}))          // Output: [1, 0]
	fmt.Println(plusOne([]int{9, 9, 9}))    // Output: [1, 0,
}
