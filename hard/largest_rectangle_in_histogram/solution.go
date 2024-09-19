package main

import (
	"fmt"
)

func largestRectangleArea(heights []int) int {
	stack := []int{}
	maxArea := 0

	for i := 0; i <= len(heights); i++ {
		var h int
		if i == len(heights) {
			h = 0 // Sentinel value to pop all remaining bars
		} else {
			h = heights[i]
		}

		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1] // Pop the top
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			maxArea = max(maxArea, height*width)
		}
		stack = append(stack, i)
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	heights1 := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights1)) // Output: 10

	heights2 := []int{2, 4}
	fmt.Println(largestRectangleArea(heights2)) // Output: 4
}
