package main

import "fmt"

func increasingTriplet(nums []int) bool {
	size := len(nums)
	for i := 2; i < size; i++ {
		if nums[i-2] < nums[i-1] && nums[i-1] < nums[i] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}
