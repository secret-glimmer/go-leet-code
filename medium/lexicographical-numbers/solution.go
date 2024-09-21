package main

import (
	"fmt"
)

func dfs(res *[]int, cur int, n int) {
	if cur > n {
		return
	}
	*res = append(*res, cur)
	for i := 0; i < 10; i++ {
		dfs(res, cur*10+i, n)
	}
}

func lexicalOrder(n int) []int {
	res := []int{}
	for i := 1; i < 10; i++ {
		dfs(&res, i, n)
	}
	return res
}

func main() {
	n := 23423
	result := lexicalOrder(n)
	fmt.Println(result) // Output: [1,10,11,12,13,2,3,4,5,6,7,8,9]
}
