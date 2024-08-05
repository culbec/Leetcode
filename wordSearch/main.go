package main

import "fmt"

func reverseWord(word string) string {
	result := ""
	n := len(word)

	for i := 0; i < n; i++ {
		result += string(word[n-i-1])
	}

	return result
}

func backtrack(board [][]byte, word string, k, i, j int) int {
	// Base cases

	// 1. Word index out of bounds.
	if k == len(word) {
		return 1
	}

	// 2. Board index out of bounds.
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == '#' {
		return 0
	}

	// 3. Character on the board not equal to the character in the word.
	if board[i][j] != word[k] {
		return 0
	}

	ch := board[i][j]
	board[i][j] = '#'

	result := backtrack(board, word, k+1, i+1, j)
	result |= backtrack(board, word, k+1, i, j+1)
	result |= backtrack(board, word, k+1, i-1, j)
	result |= backtrack(board, word, k+1, i, j-1)

	board[i][j] = ch

	return result
}

func exist(board [][]byte, word string) bool {
	freqs := make(map[byte]int)

	// Counting the frequencies of the characters.
	for _, row := range board {
		for _, el := range row {
			freqs[el]++
		}
	}

	// Verifying if all the characters in the word are present in the board.
	for _, ch := range word {
		if freqs[byte(ch)] == 0 {
			return false
		}
	}

	// Checking if we can reverse the string so that we do not have to make that many backtracks.
	firstLetter, lastLetter := word[0], word[len(word)-1]

	if freqs[firstLetter] > freqs[lastLetter] {
		word = reverseWord(word)
	}

	// Backtracking to find the word on the board.

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtrack(board, word, 0, i, j) == 1 {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED"))
}
