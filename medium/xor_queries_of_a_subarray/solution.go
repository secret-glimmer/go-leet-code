package main

import "fmt"

// Function to compute XOR for each query
func xorQueries(arr []int, queries [][]int) []int {
	// Precompute the prefix XOR array
	n := len(arr)
	prefixXor := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixXor[i+1] = prefixXor[i] ^ arr[i]
	}

	// Prepare the answer array
	answer := make([]int, len(queries))
	for i, query := range queries {
		left, right := query[0], query[1]
		answer[i] = prefixXor[right+1] ^ prefixXor[left]
	}

	return answer
}

func main() {
	// Example usage
	arr1 := []int{1, 3, 4, 8}
	queries1 := [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}
	fmt.Println(xorQueries(arr1, queries1)) // Output: [2, 7, 14, 8]

	arr2 := []int{4, 8, 2, 10}
	queries2 := [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}
	fmt.Println(xorQueries(arr2, queries2)) // Output: [8, 0, 4, 4]
}
