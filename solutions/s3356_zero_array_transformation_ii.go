/*
	https://leetcode.com/problems/zero-array-transformation-ii/description/

	You are given an integer array nums of length n and a 2D array queries where queries[i] = [li, ri, vali].

	Each queries[i] represents the following action on nums:

		Decrement the value at each index in the range [li, ri] in nums by at most vali.
		The amount by which each value is decremented can be chosen independently for each index.

	A Zero Array is an array with all its elements equal to 0.

	Return the minimum possible non-negative value of k, such that after processing the first k queries in sequence,
	nums becomes a Zero Array. If no such k exists, return -1.
*/

package solutions

func minZeroArray(nums []int, queries [][]int) int {
	l, r := 0, len(queries)
	if !canBeZero(nums, queries, r) {
		return -1
	}

	for l <= r {
		mid := (l + r) >> 1
		if canBeZero(nums, queries, mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

func canBeZero(nums []int, queries [][]int, k int) bool {
	n, sum := len(nums), 0
	diff := make([]int, n+1)

	for i := 0; i < k; i++ {
		l, r, val := queries[i][0], queries[i][1], queries[i][2]
		diff[l] += val
		diff[r+1] -= val
	}

	for i := 0; i < n; i++ {
		sum += diff[i]
		if sum < nums[i] {
			return false
		}
	}

	return true
}
