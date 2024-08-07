package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func main() {
	nums1 := []int{1, 3, 5, 6}
	target1 := 5
	fmt.Println(searchInsert(nums1, target1)) // Output: 2

	nums2 := []int{1, 3, 5, 6}
	target2 := 2
	fmt.Println(searchInsert(nums2, target2)) // Output: 1

	nums3 := []int{1, 3, 5, 6}
	target3 := 7
	fmt.Println(searchInsert(nums3, target3)) // Output: 4
}
