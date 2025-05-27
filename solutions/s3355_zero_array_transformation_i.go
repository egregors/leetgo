/*
	https://leetcode.com/problems/zero-array-transformation-i/

		You are given an integer array nums of length n and a 2D array queries, where
			queries[i] = [li, ri].

	For each queries[i]:

		Select a

		of indices within the range [li, ri] in nums.
		Decrement the values at the selected indices by 1.

	A Zero Array is an array where all elements are equal to 0.

	Return true if it is possible to transform nums into a Zero Array after
		processing all the queries sequentially,
	otherwise return false.
*/

package solutions

func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	delta := make([]int, n+1)
	for _, q := range queries {
		l, r := q[0], q[1]
		delta[l]++
		delta[r+1]--
	}

	diff := make([]int, len(delta))
	curr := 0
	for i, d := range delta {
		curr += d
		diff[i] = curr
	}

	for i := 0; i < n; i++ {
		if diff[i] < nums[i] {
			return false
		}
	}

	return true
}
