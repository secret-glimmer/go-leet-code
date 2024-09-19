package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(expression string) []int {
	if isDigit(expression) {
		num, _ := strconv.Atoi(expression)
		return []int{num}
	}

	var results []int
	for i := 0; i < len(expression); i++ {
		if isOperator(expression[i]) {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])

			for _, l := range left {
				for _, r := range right {
					results = append(results, compute(l, r, string(expression[i])))
				}
			}
		}
	}
	return results
}

func isOperator(c byte) bool {
	return c == '+' || c == '-' || c == '*'
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func compute(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	}
	return 0
}

func main() {
	expressions := []string{"2-1-1", "2*3-4*5"}
	for _, expr := range expressions {
		results := diffWaysToCompute(expr)
		fmt.Printf("Input: %s\nOutput: %v\n", expr, results)
	}
}
