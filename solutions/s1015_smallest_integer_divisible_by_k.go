/*
	https://leetcode.com/problems/smallest-integer-divisible-by-k/

	Given a positive integer k, you need to find the length of the smallest
		positive integer n such that n is divisible by k, and n only contains the
		digit 1.

	Return the length of n. If there is no such n, return -1.

	Note: n may not fit in a 64-bit signed integer.
*/

package solutions

func smallestRepunitDivByK(k int) int {
	rem := 0
	for n := 1; n <= k; n++ {
		rem = (rem*10 + 1) % k
		if rem == 0 {
			return n
		}
	}
	return -1
}
