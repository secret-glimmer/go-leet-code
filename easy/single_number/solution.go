package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {
	nums1 := []int{2, 2, 1}
	fmt.Println(singleNumber(nums1)) // Output: 1

	nums2 := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums2)) // Output: 4

	nums3 := []int{1}
	fmt.Println(singleNumber(nums3)) // Output: 1
}
