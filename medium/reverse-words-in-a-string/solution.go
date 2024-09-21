package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	result := ""
	for _, word := range words {
		if word != "" {
			result = " " + word + result
		}
	}
	return result[1:]
}

func main() {
	fmt.Println(reverseWords("the sky is blue") + "*")
	fmt.Println(reverseWords("  hello world  ") + "*")
	fmt.Println(reverseWords("a good   example") + "*")
}
