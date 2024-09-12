package main

import (
	"fmt"
)

func countConsistentStrings(allowed string, words []string) int {
	allowedSet := make(map[rune]struct{})
	for _, char := range allowed {
		allowedSet[char] = struct{}{}
	}

	count := 0
	for _, word := range words {
		isConsistent := true
		for _, char := range word {
			if _, exists := allowedSet[char]; !exists {
				isConsistent = false
				break
			}
		}
		if isConsistent {
			count++
		}
	}

	return count
}

func main() {
	allowed1 := "ab"
	words1 := []string{"ad", "bd", "aaab", "baa", "badab"}
	fmt.Println(countConsistentStrings(allowed1, words1)) // Output: 2

	allowed2 := "abc"
	words2 := []string{"a", "b", "c", "ab", "ac", "bc", "abc"}
	fmt.Println(countConsistentStrings(allowed2, words2)) // Output: 7

	allowed3 := "cad"
	words3 := []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}
	fmt.Println(countConsistentStrings(allowed3, words3)) // Output: 4
}
