/*
	https://leetcode.com/problems/smallest-number-with-all-set-bits/

	You are given a positive number n.

	Return the smallest number x greater than or equal to n, such that the binary
		representation of x contains only
*/

package solutions

func smallestNumber3370(n int) int {
	x := 1
	for x < n {
		x = x*2 + 1
	}
	return x
}
