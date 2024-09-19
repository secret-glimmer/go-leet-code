package main

import (
	"fmt"
)

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	if k >= n/2 {
		return maxProfitUnlimited(prices)
	}

	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 1; i <= k; i++ {
		maxDiff := -prices[0]
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i][j-1], prices[j]+maxDiff)
			maxDiff = max(maxDiff, dp[i-1][j]-prices[j])
		}
	}

	return dp[k][n-1]
}

func maxProfitUnlimited(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	k1, prices1 := 2, []int{2, 4, 1}
	fmt.Println(maxProfit(k1, prices1)) // Output: 2

	k2, prices2 := 2, []int{3, 2, 6, 5, 0, 3}
	fmt.Println(maxProfit(k2, prices2)) // Output: 7
}
