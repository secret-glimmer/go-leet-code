package main

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) string {
	var result strings.Builder
	i, j, carry := len(a)-1, len(b)-1, 0

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		result.WriteByte(byte(sum%2 + '0')) // Append the binary result
		carry = sum / 2                     // Update carry
	}

	// Reverse the result since we built it backwards
	resStr := reverse(result.String())
	return resStr
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// Test cases
	fmt.Println(addBinary("11", "1"))      // Output: "100"
	fmt.Println(addBinary("1010", "1011")) // Output: "10101"
	fmt.Println(addBinary("0", "0"))       // Output: "0"
	fmt.Println(addBinary("111", "111"))   // Output: "1110"
}
