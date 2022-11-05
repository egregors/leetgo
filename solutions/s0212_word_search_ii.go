/*
	https://leetcode.com/problems/word-search-ii/

	Given an m x n board of characters and a list of strings words, return all words
	on the board.

	Each word must be constructed from letters of sequentially adjacent cells, where
	adjacent cells are horizontally or vertically neighboring. The same letter cell
	may not be used more than once in a word.
*/

package solutions

func findWords(board [][]byte, words []string) []string {
	t := NewTrie[string]()
	for _, w := range words {
		t.Insert(w, w)
	}
	var res []string
	dirs := [][]int{
		{-1, 0},
		{0, -1}, {0, 1},
		{1, 0},
	}

	var bt func(r, c int, node *Trie[string])
	bt = func(r, c int, node *Trie[string]) {
		letter := board[r][c]
		if letter == '#' {
			return
		}

		if curNode := node.children[letter-'a']; curNode != nil {
			if curNode.isKey {
				res = append(res, curNode.Val)
				curNode.isKey = false
			}

			board[r][c] = '#'
			for _, d := range dirs {
				newR, newC := r+d[0], c+d[1]
				if newR >= 0 && newR < len(board) && newC >= 0 && newC < len(board[0]) {
					bt(newR, newC, curNode)
				}
			}
			board[r][c] = letter
		}
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			bt(r, c, t)
		}
	}

	return res
}
