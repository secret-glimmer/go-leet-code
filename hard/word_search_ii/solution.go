package main

import (
	"fmt"
)

func findWords(board [][]byte, words []string) []string {
	trie := buildTrie(words)
	result := []string{}
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			backtrack(board, i, j, trie, &result, "", visited)
		}
	}

	return result
}

func backtrack(board [][]byte, i, j int, trie *TrieNode, result *[]string, currentWord string, visited [][]bool) {
	if visited[i][j] {
		return
	}

	currentWord += string(board[i][j])
	if node, exists := trie.children[board[i][j]]; exists {
		if node.isEnd {
			*result = append(*result, currentWord)
			node.isEnd = false // prevent duplicates
		}

		visited[i][j] = true
		directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, dir := range directions {
			newI, newJ := i+dir[0], j+dir[1]
			if newI >= 0 && newI < len(board) && newJ >= 0 && newJ < len(board[0]) {
				backtrack(board, newI, newJ, node, result, currentWord, visited)
			}
		}
		visited[i][j] = false
	}
}

type TrieNode struct {
	children map[byte]*TrieNode
	isEnd    bool
}

func buildTrie(words []string) *TrieNode {
	root := &TrieNode{children: make(map[byte]*TrieNode)}
	for _, word := range words {
		node := root
		for i := 0; i < len(word); i++ {
			if _, exists := node.children[word[i]]; !exists {
				node.children[word[i]] = &TrieNode{children: make(map[byte]*TrieNode)}
			}
			node = node.children[word[i]]
		}
		node.isEnd = true
	}
	return root
}

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(board, words)) // Output: ["eat", "oath"]
}
