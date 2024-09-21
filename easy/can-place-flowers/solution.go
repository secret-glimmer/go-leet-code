package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	startIndex := -1
	for index, isExisted := range flowerbed {
		if isExisted == 1 && startIndex == -1 {
			startIndex = index
		} else if isExisted == 1 {
			if (index-startIndex)%(n+1) != 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 1, 0, 1, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}
