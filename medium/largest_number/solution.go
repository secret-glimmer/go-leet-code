package main

import (
	"fmt"
	"sort"
	"strings"
)

// Custom type for sorting
type ByConcat []string

// Implementing the sort.Interface for ByConcat
func (a ByConcat) Len() int      { return len(a) }
func (a ByConcat) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByConcat) Less(i, j int) bool {
	// Compare concatenated strings
	return a[i]+a[j] > a[j]+a[i]
}

func largestNumber(nums []int) string {
	// Convert integers to strings
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = fmt.Sprintf("%d", num)
	}

	// Sort strings using custom comparator
	sort.Sort(ByConcat(strs))

	// If the largest number is "0", return "0"
	if strs[0] == "0" {
		return "0"
	}

	// Join sorted strings to form the largest number
	return strings.Join(strs, "")
}

func main() {
	// Test cases
	nums1 := []int{10, 2}
	nums2 := []int{3, 30, 34, 5, 9}

	fmt.Println(largestNumber(nums1)) // Output: "210"
	fmt.Println(largestNumber(nums2)) // Output: "9534330"
}
