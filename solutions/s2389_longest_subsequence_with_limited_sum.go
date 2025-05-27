/*
	2389. Longest Subsequence With Limited Sum


	You are given an integer array nums of length n, and an integer array queries
		of length m.

	Return an array answer of length m where answer[i] is the maximum size of a
		subsequence
	that you can take from nums such that the sum of its elements is less than or
		equal to
	queries[i].

	A subsequence is an array that can be derived from another array by deleting
		some or no
	elements without changing the order of the remaining elements.
*/

package solutions

import "sort"

func answerQueries(nums, queries []int) []int {
	sort.Ints(nums)
	res := make([]int, len(queries))
	for i, target := range queries {
		var sum, length int
		for _, n := range nums {
			if sum+n <= target {
				sum += n
				length++
			}
		}
		res[i] = length
	}
	return res
}
