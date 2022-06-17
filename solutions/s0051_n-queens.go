/*
	https://leetcode.com/problems/n-queens/

	The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no
	two queens attack each other.

	Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the
	answer in any order.

	Each solution contains a distinct board configuration of the n-queens' placement, where 'Q'
	and '.' both indicate a queen and an empty space, respectively.
*/

//nolint:revive // it's ok
package solutions

func solveNQueens(n int) [][]string {
	emptyBoard := make([][]rune, n)
	for i := 0; i < n; i++ {
		emptyBoard[i] = make([]rune, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			emptyBoard[i][j] = '.'
		}
	}

	var res [][]string
	backtrack51(0, n, make(Set[int]), make(Set[int]), make(Set[int]), emptyBoard, &res)
	return res
}

func backtrack51(row, n int, diagonals, antiDiagonals, cols Set[int], state [][]rune, acc *[][]string) {
	if row == n {
		var rs []string
		for _, r := range state {
			rs = append(rs, string(r))
		}
		*acc = append(*acc, rs)
		return
	}

	for col := 0; col < n; col++ {
		currDiagonal := row - col
		currAntiDiagonal := row + col
		if cols.Contains(col) || diagonals.Contains(currDiagonal) || antiDiagonals.Contains(currAntiDiagonal) {
			continue
		}

		cols.Add(col)
		diagonals.Add(currDiagonal)
		antiDiagonals.Add(currAntiDiagonal)
		state[row][col] = 'Q'

		backtrack51(row+1, n, diagonals, antiDiagonals, cols, state, acc)

		cols.Remove(col)
		diagonals.Remove(currDiagonal)
		antiDiagonals.Remove(currAntiDiagonal)
		state[row][col] = '.'
	}
}
