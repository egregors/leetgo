/*
	https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-i/

	You are given an integer array nums.

	A subsequence sub of nums with length x is called valid if it satisfies:

		(sub[0] + sub[1]) % 2 == (sub[1] + sub[2]) % 2 == ... == (sub[x - 2] + sub[x -
			1]) % 2.

	Return the length of the longest valid subsequence of nums.

	A subsequence is an array that can be derived from another array by deleting
		some or no elements without changing the order of the remaining elements.
*/

package solutions

func maximumLength(nums []int) int {
	res := 0
	patterns := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	for _, pattern := range patterns {
		cnt := 0
		for _, num := range nums {
			if num%2 == pattern[cnt%2] {
				cnt++
			}
		}
		res = max(res, cnt)
	}
	return res
}
