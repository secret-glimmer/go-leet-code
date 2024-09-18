package main

import (
	"fmt"
)

func maxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		// Calculate the width and height
		width := right - left
		h := min(height[left], height[right])
		// Calculate the area
		area := width * h
		// Update maxArea if we found a larger area
		if area > maxArea {
			maxArea = area
		}

		// Move the pointer pointing to the shorter line
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height1)) // Output: 49

	height2 := []int{1, 1}
	fmt.Println(maxArea(height2)) // Output: 1
}
