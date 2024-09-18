package main

import (
	"fmt"
)

// Function to compute the GCD of two numbers
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Function to check if str can be formed by repeating base
func canFormByRepeating(str, base string) bool {
	if len(str)%len(base) != 0 {
		return false
	}
	repeatCount := len(str) / len(base)
	repeated := ""
	for i := 0; i < repeatCount; i++ {
		repeated += base
	}
	return repeated == str
}

// Function to find the greatest common divisor of two strings
func gcdOfStrings(str1 string, str2 string) string {
	// Get the lengths of the strings
	len1 := len(str1)
	len2 := len(str2)

	// Find the GCD of the lengths
	gcdLength := gcd(len1, len2)

	// Extract the potential common divisor string
	potentialGCD := str1[:gcdLength]

	// Check if it can form both strings
	if canFormByRepeating(str1, potentialGCD) && canFormByRepeating(str2, potentialGCD) {
		return potentialGCD
	}

	return ""
}

func main() {
	// Test cases
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))  // Output: "ABC"
	fmt.Println(gcdOfStrings("ABABAB", "ABAB")) // Output: "AB"
	fmt.Println(gcdOfStrings("LEET", "CODE"))   // Output: ""
}
