/*
	https://leetcode.com/problems/snakes-and-ladders/

	You are given an n x n integer matrix board where the cells are labeled from 1 to n2 in a
	Boustrophedon style starting from the bottom left of the board (i.e. board[n - 1][0]) and
	alternating direction each row.

	You start on square 1 of the board. In each move, starting from square curr, do the
	following:

		Choose a destination square next with a label in the range [curr + 1,
	min(curr + 6, n2)].
			This choice simulates the result of a standard 6-sided die roll: i.e., there
	are always at most 6 destinations, regardless of the size of the board.
		If next has a snake or ladder, you must move to the destination of that snake or
	ladder. Otherwise, you move to next.
		The game ends when you reach the square n2.

	A board square on row r and column c has a snake or ladder if board[r][c] != -1.
	The destination of that snake or ladder is board[r][c]. Squares 1 and n2 do not have
	a snake or ladder.

	Note that you only take a snake or ladder at most once per move. If the destination
	to a snake or ladder is the start of another snake or ladder, you do not follow the
	subsequent snake or ladder.

		For example, suppose the board is [[-1,4],[-1,3]], and on the first move, your
	destination square is 2. You follow the ladder to square 3, but do not follow the
	subsequent ladder to 4.

	Return the least number of moves required to reach the square n2. If it is not possible
	to reach the square, return -1.
*/

package solutions

func snakesAndLadders(board [][]int) int {
	getIndex := func(x int, n int) (i int, j int) {
		x--
		i, j = x/n, x%n
		if i%2 == 1 {
			j = n - 1 - j
		}
		i = n - 1 - i
		return
	}

	queue, moves := []int{1}, 0
	n := len(board)
	visit := map[int]bool{1: true}
	for len(queue) > 0 {
		size := len(queue)
		for k := 0; k < size; k++ {
			x := queue[k]
			if x == n*n {
				return moves
			}
			for l := 1; l <= 6; l++ {
				x++
				if x > n*n {
					break
				}
				i, j := getIndex(x, n)
				target := x
				if board[i][j] != -1 {
					target = board[i][j]
				}
				if !visit[target] {
					queue = append(queue, target)
					visit[target] = true
				}
			}
		}
		queue = queue[size:]
		moves++
	}
	return -1
}
