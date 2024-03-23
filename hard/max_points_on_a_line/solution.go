package main

import (
	"fmt"
	"math"
)

func maxPoints(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}

	maxPoints := 0

	for i := 0; i < len(points); i++ {
		slopeCount := make(map[float64]int)
		overlap := 0
		currentMax := 0

		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}

			dx := float64(points[j][0] - points[i][0])
			dy := float64(points[j][1] - points[i][1])

			if dx == 0 && dy == 0 {
				overlap++
				continue
			}

			slope := math.Atan2(dy, dx)
			slopeCount[slope]++

			if slopeCount[slope] > currentMax {
				currentMax = slopeCount[slope]
			}
		}

		maxPoints = max(maxPoints, currentMax+overlap+1) // +1 for the point itself
	}

	return maxPoints
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	points1 := [][]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println(maxPoints(points1)) // Output: 3

	points2 := [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}
	fmt.Println(maxPoints(points2)) // Output: 4
}
