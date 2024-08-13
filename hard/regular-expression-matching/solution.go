package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	// Create a 2D array for dynamic programming
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// Base case: empty string and empty pattern match
	dp[0][0] = true

	// Handle patterns like a*, a*b*, a*b*c* that can match an empty string
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1] // Match single character or '.'
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			}
		}
	}

	return dp[m][n]
}

func main() {
	fmt.Println(isMatch("aa", "a"))                   // Output: false
	fmt.Println(isMatch("aa", "a*"))                  // Output: true
	fmt.Println(isMatch("ab", ".*"))                  // Output: true
	fmt.Println(isMatch("aab", "c*a*b"))              // Output: true
	fmt.Println(isMatch("mississippi", "mis*is*p*.")) // Output: false
}
