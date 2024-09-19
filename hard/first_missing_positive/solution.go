package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	n := len(nums)

	// Place each number in its right place
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			// Swap nums[i] with nums[nums[i]-1]
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	// Find the first missing positive
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

func main() {
	nums1 := []int{1, 2, 0}
	fmt.Println(firstMissingPositive(nums1)) // Output: 3

	nums2 := []int{3, 4, -1, 1}
	fmt.Println(firstMissingPositive(nums2)) // Output: 2

	nums3 := []int{7, 8, 9, 11, 12}
	fmt.Println(firstMissingPositive(nums3)) // Output: 1
}
