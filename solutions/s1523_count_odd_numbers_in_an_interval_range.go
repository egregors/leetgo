/*
	https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/

	Given two non-negative integers low and high. Return the count of odd numbers between
	low and high (inclusive).
*/

package solutions

func countOdds(low, high int) int {
	if low%2 != 0 || high%2 != 0 {
		return (high-low)/2 + 1
	}
	return (high - low) / 2
}
