/*
	https://leetcode.com/problems/power-of-three/

	Given an integer n, return true if it is a power of three. Otherwise, return
		false.

	An integer n is a power of three, if there exists an integer x such that n ==
		3x.
*/

package solutions

import (
	"math"
)

func isPowerOfThree(n int) bool {
	m := 1
	for m < math.MaxInt32/2 {
		m *= 3
	}
	return n > 0 && m%n == 0
}
