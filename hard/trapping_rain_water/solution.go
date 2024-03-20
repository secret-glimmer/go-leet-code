package main

import (
	"fmt"
)

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	// Fill leftMax array
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	// Fill rightMax array
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	// Calculate trapped water
	trappedWater := 0
	for i := 0; i < n; i++ {
		trappedWater += min(leftMax[i], rightMax[i]) - height[i]
	}

	return trappedWater
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height1)) // Output: 6

	height2 := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height2)) // Output: 9
}
