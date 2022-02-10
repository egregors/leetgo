/*
	https://leetcode.com/problems/word-search/

	Given an m x n grid of characters board and a string word, return true if word exists
	in the grid.

	The word can be constructed from letters of sequentially adjacent cells, where adjacent
	cells are horizontally or vertically neighboring. The same letter cell may not be used
	more than once.
*/

package solutions

import "bytes"

func backtrack79(r, c int, curr, word []byte, seen map[[2]int]bool, board [][]byte) bool {
	if bytes.Equal(curr, word) {
		return true
	}

	dirs := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

	for _, d := range dirs {
		nextR, nextC := r+d[0], c+d[1]
		if nextR >= 0 && nextR < len(board) && nextC >= 0 && nextC < len(board[0]) {
			if seen[[2]int{nextR, nextC}] || board[nextR][nextC] != word[len(curr)] {
				continue
			}
			seen[[2]int{nextR, nextC}] = true
			curr = append(curr, board[nextR][nextC])
			if backtrack79(nextR, nextC, curr, word, seen, board) {
				return true
			}
			curr = curr[:len(curr)-1]
			seen[[2]int{nextR, nextC}] = false
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	bWord := []byte(word)
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if board[r][c] == bWord[0] &&
				backtrack79(r, c, []byte{board[r][c]}, bWord, map[[2]int]bool{{r, c}: true}, board) {
				return true
			}
		}
	}
	return false
}
