/*
	https://leetcode.com/problems/surrounded-regions/

	Given an m x n matrix board containing 'X' and 'O', capture all regions that are
	4-directionally surrounded by 'X'.

	A region is captured by flipping all 'O's into 'X's in that surrounded region.
*/

package solutions

const (
	borderline = '*'
	x          = 'X'
	o          = 'O'
)

func solve(board [][]byte) {
	n, m := len(board), len(board[0])

	if n < 3 || m < 3 {
		return
	}

	f := func(r, c int) {
		dfs(r, c, n, m, board)
	}

	// left | right borders
	for r := 0; r < n; r++ {
		if board[r][0] == o {
			f(r, 0)
		}
		if board[r][m-1] == o {
			f(r, m-1)
		}
	}

	// upper | bottom borders
	for c := 1; c < m-1; c++ {
		if board[0][c] == o {
			f(0, c)
		}
		if board[n-1][c] == o {
			f(n-1, c)
		}
	}

	// whole board
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			switch board[r][c] {
			case o:
				board[r][c] = x
			case borderline:
				board[r][c] = o
			}
		}
	}

}

func dfs(r, c, n, m int, board [][]byte) {
	board[r][c] = borderline
	// up
	if r > 0 && board[r-1][c] == o {
		dfs(r-1, c, n, m, board)
	}
	// down
	if r < n-1 && board[r+1][c] == o {
		dfs(r+1, c, n, m, board)
	}
	// left
	if c > 0 && board[r][c-1] == o {
		dfs(r, c-1, n, m, board)
	}
	// right
	if c < m-1 && board[r][c+1] == o {
		dfs(r, c+1, n, m, board)
	}
}
