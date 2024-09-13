package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	// Trim any trailing spaces
	s = strings.TrimSpace(s)
	// Split the string by spaces
	words := strings.Split(s, " ")
	// Return the length of the last word
	return len(words[len(words)-1])
}

func main() {
	// Test cases
	fmt.Println(lengthOfLastWord("Hello World"))                 // Output: 5
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  ")) // Output: 4
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))       // Output: 6
}
