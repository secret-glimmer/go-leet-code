package main

import (
	"fmt"
)

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])

	// Create a dp array to store the minimum health needed at each cell
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Start from the princess's room
	dp[m-1][n-1] = max(1, 1-dungeon[m-1][n-1])

	// Fill the last row
	for j := n - 2; j >= 0; j-- {
		dp[m-1][j] = max(1, dp[m-1][j+1]-dungeon[m-1][j])
	}

	// Fill the last column
	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = max(1, dp[i+1][n-1]-dungeon[i][n-1])
	}

	// Fill the rest of the dp table
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			minHealthOnExit := min(dp[i+1][j], dp[i][j+1])
			dp[i][j] = max(1, minHealthOnExit-dungeon[i][j])
		}
	}

	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	dungeon1 := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	fmt.Println(calculateMinimumHP(dungeon1)) // Output: 7

	dungeon2 := [][]int{{0}}
	fmt.Println(calculateMinimumHP(dungeon2)) // Output: 1
}
