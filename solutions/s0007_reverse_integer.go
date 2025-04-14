/*
	https://leetcode.com/problems/reverse-integer/description/

	Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go
	outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

	Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
*/

package solutions

import "math"

func reverse0007(x int) int {
	lo, hi := math.MinInt32, math.MaxInt32
	res := 0
	for x != 0 {
		pop := x % 10
		x /= 10

		res = res*10 + pop
		if res > hi || res < lo {
			return 0
		}
	}

	return res
}
