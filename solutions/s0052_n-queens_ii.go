/*
	https://leetcode.com/problems/n-queens/

	The n-queens puzzle is the problem of placing n queens on an n x n chessboard
		such that no
	two queens attack each other.

	Given an integer n, return all distinct solutions to the n-queens puzzle. You
		may return the
	answer in any order.

	Each solution contains a distinct board configuration of the n-queens'
		placement, where 'Q'
	and '.' both indicate a queen and an empty space, respectively.
*/

//nolint:revive // it's ok
package solutions

func totalNQueens(n int) int {
	emptyBoard := make([][]rune, n)
	for i := 0; i < n; i++ {
		emptyBoard[i] = make([]rune, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			emptyBoard[i][j] = '.'
		}
	}

	return backtrack52(0, n, make(Set[int]), make(Set[int]), make(Set[int]), emptyBoard)
}

func backtrack52(row, n int, diagonals, antiDiagonals, cols Set[int], state [][]rune) int {
	if row == n {
		return 1
	}
	var res int

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

		res += backtrack52(row+1, n, diagonals, antiDiagonals, cols, state)

		cols.Remove(col)
		diagonals.Remove(currDiagonal)
		antiDiagonals.Remove(currAntiDiagonal)
		state[row][col] = '.'
	}
	return res
}
