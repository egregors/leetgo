/*
	https://leetcode.com/problems/push-dominoes/description/

	There are n dominoes in a line, and we place each domino vertically upright. In
		the beginning, we simultaneously
	push some of the dominoes either to the left or to the right.

	After each second, each domino that is falling to the left pushes the adjacent
		domino on the left. Similarly,
	the dominoes falling to the right push their adjacent dominoes standing on the
		right.

	When a vertical domino has dominoes falling on it from both sides, it stays
		still due to the balance of the forces.

	For the purposes of this question, we will consider that a falling domino
		expends no additional force to a
	falling or already fallen domino.

	You are given a string dominoes representing the initial state where:

		dominoes[i] = 'L', if the ith domino has been pushed to the left,
		dominoes[i] = 'R', if the ith domino has been pushed to the right, and
		dominoes[i] = '.', if the ith domino has not been pushed.

	Return a string representing the final state.

*/

package solutions

import "strings"

//nolint:staticcheck // meh
func pushDominoes(dominoes string) string {
	xs := []rune(dominoes)
	n := len(xs)
	forces := make([]int, n)

	f := 0
	for i := 0; i < n; i++ {
		if xs[i] == 'R' {
			f = n
		} else if xs[i] == 'L' {
			f = 0
		} else {
			f = max(f-1, 0)
		}
		forces[i] += f
	}

	f = 0
	for i := n - 1; i >= 0; i-- {
		if xs[i] == 'L' {
			f = n
		} else if xs[i] == 'R' {
			f = 0
		} else {
			f = max(f-1, 0)
		}
		forces[i] -= f
	}

	sb := strings.Builder{}
	for _, f := range forces {
		switch {
		case f > 0:
			sb.WriteRune('R')
		case f < 0:
			sb.WriteRune('L')
		case f == 0:
			sb.WriteRune('.')
		}
	}

	return sb.String()
}
