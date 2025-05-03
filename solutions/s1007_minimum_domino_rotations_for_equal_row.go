/*
	https://leetcode.com/problems/minimum-domino-rotations-for-equal-row/description/

	In a row of dominoes, tops[i] and bottoms[i] represent the top and bottom halves of the ith domino.
	(A domino is a tile with two numbers from 1 to 6 - one on each half of the tile.)

	We may rotate the ith domino, so that tops[i] and bottoms[i] swap values.

	Return the minimum number of rotations so that all the values in tops are the same, or all the values in
	bottoms are the same.

	If it cannot be done, return -1.
*/

package solutions

func minDominoRotations(tops []int, bottoms []int) int {
	check := func(x int, as, bs []int, n int) int {
		aFlips, bFlips := 0, 0
		for i := 0; i < n; i++ {
			if as[i] != x && bs[i] != x {
				return -1
			} else if as[i] != x {
				aFlips++
			} else if bs[i] != x {
				bFlips++
			}
		}
		return min(aFlips, bFlips)
	}
	n := len(tops)
	res := check(tops[0], bottoms, tops, n)
	if res != -1 || tops[0] == bottoms[0] {
		return res
	}
	return check(bottoms[0], bottoms, tops, n)
}
