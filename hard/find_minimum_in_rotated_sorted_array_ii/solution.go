package main

import (
	"fmt"
)

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		// Handle duplicates
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		for left < right && nums[right] == nums[right-1] {
			right--
		}

		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}

func main() {
	nums1 := []int{1, 3, 5}
	fmt.Println(findMin(nums1)) // Output: 1

	nums2 := []int{2, 2, 2, 0, 1}
	fmt.Println(findMin(nums2)) // Output: 0
}
