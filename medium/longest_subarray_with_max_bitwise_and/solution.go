package main

import (
	"fmt"
)

func longestSubarrayWithMaxBitwiseAND(nums []int) int {
	maxValue := 0
	longestLength := 0
	currentLength := 0

	// Find the maximum value in the array
	for _, num := range nums {
		if num > maxValue {
			maxValue = num
		}
	}

	// Iterate through the array to find the longest subarray with the maximum value
	for _, num := range nums {
		if num == maxValue {
			currentLength++
			if currentLength > longestLength {
				longestLength = currentLength
			}
		} else {
			currentLength = 0
		}
	}

	return longestLength
}

func main() {
	nums1 := []int{1, 2, 3, 3, 2, 2}
	fmt.Println(longestSubarrayWithMaxBitwiseAND(nums1)) // Output: 2

	nums2 := []int{1, 2, 3, 4}
	fmt.Println(longestSubarrayWithMaxBitwiseAND(nums2)) // Output: 1
}
