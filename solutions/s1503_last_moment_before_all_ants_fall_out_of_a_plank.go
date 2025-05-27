/*
	https://leetcode.com/problems/last-moment-before-all-ants-fall-out-of-a-plank/description/

	We have a wooden plank of the length n units. Some ants are walking on the
		plank, each ant moves with a speed of
	1 unit per second. Some of the ants move to the left, the other move to the
		right.

	When two ants moving in two different directions meet at some point, they
		change their directions and continue
	moving again. Assume changing directions does not take any additional time.

	When an ant reaches one end of the plank at a time t, it falls out of the plank
		immediately.

	Given an integer n and two integer arrays left and right, the positions of the
		ants moving to the left and the
	right, return the moment when the last ant(s) fall out of the plank.
*/

package solutions

func getLastMoment(n int, left, right []int) int {
	var res int

	for _, x := range left {
		res = max(res, x)
	}
	for _, y := range right {
		res = max(res, n-y)
	}

	return res
}
