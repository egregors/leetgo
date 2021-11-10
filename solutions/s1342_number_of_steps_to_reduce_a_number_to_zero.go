/*
	https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/

	Given an integer num, return the number of steps to reduce it to zero.

	In one step, if the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.
*/

package solutions

func numberOfSteps(num int) int {
	steps := 0
	for num != 0 {
		steps++
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
	}
	return steps
}
