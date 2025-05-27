/*
	https://leetcode.com/problems/baseball-game/

	You are keeping score for a baseball game with strange rules. The game consists
		of several
	rounds, where the scores of past rounds may affect future rounds' scores.

	At the beginning of the game, you start with an empty record. You are given a
		list of strings
	ops, where ops[i] is the ith operation you must apply to the record and is one
		of the following:

		An integer x - Record a new score of x.
		"+" - Record a new score that is the sum of the previous two scores. It is
			guaranteed
	there will always be two previous scores.
		"D" - Record a new score that is double the previous score. It is guaranteed
			there will
	always be a previous score.
		"C" - Invalidate the previous score, removing it from the record. It is
			guaranteed there
	will always be a previous score.

	Return the sum of all the scores on the record.
*/

package solutions

import "strconv"

func calPoints(ops []string) int {
	var acc []int

	for _, op := range ops {
		switch op {
		case "+":
			acc = append(acc, Sum(acc[len(acc)-1], acc[len(acc)-2]))
		case "D":
			acc = append(acc, acc[len(acc)-1]*2)
		case "C":
			acc = acc[:len(acc)-1]
		default:
			n, _ := strconv.Atoi(op)
			acc = append(acc, n)
		}
	}

	return Sum(acc...)
}
