package main

import (
	"fmt"
)

func romanToInt(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prevValue := 0

	for _, char := range s {
		currentValue := romanMap[char]
		if currentValue > prevValue {
			total += currentValue - 2*prevValue
		} else {
			total += currentValue
		}
		prevValue = currentValue
	}

	return total
}

func main() {
	fmt.Println(romanToInt("III"))     // Output: 3
	fmt.Println(romanToInt("LVIII"))   // Output: 58
	fmt.Println(romanToInt("MCMXCIV")) // Output: 1994
}
