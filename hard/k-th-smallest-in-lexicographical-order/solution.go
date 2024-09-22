package main

import "fmt"

func dfs(res *[]int, cur int, n int) {
	if cur > n {
		return
	}
	*res = append(*res, cur)
	for i := 0; i < 10; i++ {
		dfs(res, cur*10+i, n)
	}
}

func findKthNumber(n int, k int) int {
	res := []int{}
	for i := 1; i < 10; i++ {
		dfs(&res, i, n)
		if len(res) >= k {
			return res[k-1]
		}
	}
	return res[k-1]
}

func main() {
	fmt.Println(findKthNumber(13, 3))
}
