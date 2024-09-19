package main

import (
	"fmt"
)

func solveNQueens(n int) [][]string {
	var results [][]string
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	placeQueens(board, 0, n, &results)
	return results
}

func placeQueens(board [][]byte, row, n int, results *[][]string) {
	if row == n {
		var solution []string
		for _, b := range board {
			solution = append(solution, string(b))
		}
		*results = append(*results, solution)
		return
	}

	for col := 0; col < n; col++ {
		if isSafe(board, row, col, n) {
			board[row][col] = 'Q'
			placeQueens(board, row+1, n, results)
			board[row][col] = '.' // backtrack
		}
	}
}

func isSafe(board [][]byte, row, col, n int) bool {
	// Check the column
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// Check the upper left diagonal
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// Check the upper right diagonal
	for i, j := row, col; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

func main() {
	n1 := 4
	solutions1 := solveNQueens(n1)
	fmt.Println(solutions1) // Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]

	n2 := 1
	solutions2 := solveNQueens(n2)
	fmt.Println(solutions2) // Output: [["Q"]]
}
