package main

import "fmt"

func productExceptSelf(nums []int) []int {
	temp := 1
	zero := 0
	for _, num := range nums {
		if num == 0 {
			zero++
		} else {
			temp *= num
		}
	}
	for i, num := range nums {
		if zero > 1 {
			nums[i] = 0
		} else if num == 0 && zero == 1 {
			nums[i] = temp
		} else if num != 0 && zero == 1 {
			nums[i] = 0
		} else {
			nums[i] = temp / num
		}
	}
	return nums
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
