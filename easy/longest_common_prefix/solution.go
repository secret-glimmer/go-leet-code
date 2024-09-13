package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	temp := strs[0]
	n := len(temp)
	for _, str := range strs[1:] {
		i := 0
		m := len(str)
		for i < n && i < m && str[i] == temp[i] {
			i++
		}
		n = i
	}

	return temp[:n]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // Output: 3
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // Output: 58
}
