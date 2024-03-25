package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) []string {
	wordSet := make(map[string]struct{})
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}

	memo := make(map[string][]string)
	return backtrack(s, wordSet, memo)
}

func backtrack(s string, wordSet map[string]struct{}, memo map[string][]string) []string {
	if val, found := memo[s]; found {
		return val
	}

	if len(s) == 0 {
		return []string{""}
	}

	var sentences []string
	for i := 1; i <= len(s); i++ {
		word := s[:i]
		if _, exists := wordSet[word]; exists {
			subSentences := backtrack(s[i:], wordSet, memo)
			for _, subSentence := range subSentences {
				if subSentence == "" {
					sentences = append(sentences, word)
				} else {
					sentences = append(sentences, word+" "+subSentence)
				}
			}
		}
	}

	memo[s] = sentences
	return sentences
}

func main() {
	s1 := "catsanddog"
	wordDict1 := []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println(wordBreak(s1, wordDict1)) // Output: ["cats and dog", "cat sand dog"]

	s2 := "pineapplepenapple"
	wordDict2 := []string{"apple", "pen", "applepen", "pine", "pineapple"}
	fmt.Println(wordBreak(s2, wordDict2)) // Output: ["pine apple pen apple", "pineapple pen apple", "pine applepen apple"]

	s3 := "catsandog"
	wordDict3 := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s3, wordDict3)) // Output: []
}
