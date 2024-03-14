package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	firstBuy := -prices[0]
	firstSell := 0
	secondBuy := -prices[0]
	secondSell := 0

	for _, price := range prices {
		firstBuy = max(firstBuy, -price)              // Maximize the first buy
		firstSell = max(firstSell, firstBuy+price)    // Maximize first sell
		secondBuy = max(secondBuy, firstSell-price)   // Maximize second buy
		secondSell = max(secondSell, secondBuy+price) // Maximize second sell
	}

	return secondSell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices1 := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(prices1)) // Output: 6

	prices2 := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices2)) // Output: 4

	prices3 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices3)) // Output: 0
}
