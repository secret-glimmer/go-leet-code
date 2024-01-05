package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {
	haystack1 := "sadbutsad"
	needle1 := "sad"
	fmt.Println(strStr(haystack1, needle1)) // Output: 0

	haystack2 := "leetcode"
	needle2 := "leeto"
	fmt.Println(strStr(haystack2, needle2)) // Output: -1
}
