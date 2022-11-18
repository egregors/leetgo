/*
	https://leetcode.com/problems/ugly-number/

	An ugly number is a positive integer whose prime factors are
	limited to 2, 3, and 5.

	Given an integer n, return true if n is an ugly number.
*/

package solutions

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	div := func(n, d int) int {
		for n%d == 0 {
			n /= d
		}
		return n
	}
	for _, d := range []int{2, 3, 5} {
		n = div(n, d)
	}
	return n == 1
}
