/*
	https://leetcode.com/problems/the-two-sneaky-numbers-of-digitville/

	In the town of Digitville, there was a list of numbers called nums containing
		integers from 0 to n - 1. Each number was supposed to appear exactly once in
		the list, however, two mischievous numbers sneaked in an additional time,
		making the list longer than usual.

	As the town detective, your task is to find these two sneaky numbers. Return an
		array of size two containing the two numbers (in any order), so peace can
		return to Digitville.
*/

package solutions

func getSneakyNumbers(nums []int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	var res []int
	for k, v := range m {
		if v == 2 {
			res = append(res, k)
		}
		if len(res) == 2 {
			return res
		}
	}

	return res
}
