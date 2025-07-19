/*
	https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts

		You are given a non-negative integer array nums. In one operation, you must:

		Choose a positive integer x such that x is less than or equal to the smallest
			non-zero element in nums.
		Subtract x from every positive element in nums.

	Return the minimum number of operations to make every element in nums equal to
		0.
*/

package solutions

func minimumOperations2357(nums []int) int {
	cnt := 0
	for !func(xs []int) bool {
		for _, x := range xs {
			if x > 0 {
				return false
			}
		}
		return true
	}(nums) {
		k := 101
		for _, n := range nums {
			if n != 0 {
				k = min(k, n)
			}
		}
		for i := 0; i < len(nums); i++ {
			if nums[i] > 0 {
				nums[i] -= k
			}
		}
		cnt++
	}
	return cnt
}
