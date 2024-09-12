package main

import "fmt"

func isHappy(n int) bool {
	seen := make(map[int]bool)

	for n != 1 {
		if seen[n] {
			return false
		}
		seen[n] = true
		n = sumOfSquares(n)
	}
	return true
}

func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println(isHappy(19)) // Output: true
	fmt.Println(isHappy(2))  // Output: false
}
