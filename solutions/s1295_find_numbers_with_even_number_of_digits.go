/*
	https://leetcode.com/problems/find-numbers-with-even-number-of-digits/editorial

	Given an array nums of integers, return how many of them contain an even number
		of digits.
*/

package solutions

func findNumbers(nums []int) int {
	c := 0

	pred := func(x int) bool {
		cnt := 0
		for x != 0 {
			cnt++
			x /= 10
		}

		return cnt%2 == 0
	}

	for _, n := range nums {
		if pred(n) {
			c++
		}
	}

	return c
}
