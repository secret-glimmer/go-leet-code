package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	halfLen := (m + n + 1) / 2
	minIdx, maxIdx := 0, m
	var maxOfLeft, minOf_right float64

	for minIdx <= maxIdx {
		i := (minIdx + maxIdx) / 2
		j := halfLen - i

		if i < m && nums2[j-1] > nums1[i] {
			minIdx = i + 1 // i is too small
		} else if i > 0 && nums1[i-1] > nums2[j] {
			maxIdx = i - 1 // i is too big
		} else { // i is perfect
			if i == 0 {
				maxOfLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxOfLeft = float64(nums1[i-1])
			} else {
				maxOfLeft = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}

			if (m+n)%2 == 1 {
				return maxOfLeft // Odd case
			}

			if i == m {
				minOf_right = float64(nums2[j])
			} else if j == n {
				minOf_right = float64(nums1[i])
			} else {
				minOf_right = math.Min(float64(nums1[i]), float64(nums2[j]))
			}

			return (maxOfLeft + minOf_right) / 2.0 // Even case
		}
	}
	return 0.0
}
func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) // Output: 2.00000

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) // Output: 2.50000
}
